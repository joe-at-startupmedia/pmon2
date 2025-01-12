package repo

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"pmon3/model"
	"pmon3/pmond"
	"pmon3/pmond/db"
	"sync"
)

type ProcessRepo struct {
	db  *gorm.DB
	cur *model.Process
}

var processOnce sync.Once

func Process() *ProcessRepo {
	dbInst := db.Db()
	processOnce.Do(func() {
		if !dbInst.Migrator().HasTable(&model.Process{}) {
			dbInst.Migrator().CreateTable(&model.Process{})
		}
	})
	return &ProcessRepo{
		db: dbInst,
	}
}

func ProcessOf(p *model.Process) *ProcessRepo {
	pr := Process()
	pr.cur = p
	return pr
}

func (pr *ProcessRepo) getDb() *gorm.DB {
	return pr.db.Preload("Groups")
}

func (pr *ProcessRepo) Save() error {
	return pr.db.Save(&pr.cur).Error
}

func (pr *ProcessRepo) Delete() error {
	return pr.db.Select(clause.Associations).Delete(&pr.cur).Error
}

func (pr *ProcessRepo) UpdateStatus(status model.ProcessStatus) error {
	pr.cur.Status = status
	return pr.Save()
}

func (pr *ProcessRepo) FindById(id uint32) (*model.Process, error) {
	var found model.Process
	err := pr.getDb().First(&found, id).Error
	if err != nil {
		pmond.Log.Infof("could not find process in database: %d %-v", id, err)
		return nil, err
	}
	return &found, nil
}

func (pr *ProcessRepo) FindByIdOrName(idOrName string) (*model.Process, error) {
	var found model.Process
	err := pr.getDb().First(&found, "id = ? or name = ?", idOrName, idOrName).Error
	if err != nil {
		pmond.Log.Infof("could not find process in database: %s %-v", idOrName, err)
		return nil, err
	}
	return &found, nil
}

func (pr *ProcessRepo) FindByFileAndName(processFile string, name string) (*model.Process, error) {
	var found model.Process
	err := pr.getDb().First(&found, "process_file = ? AND name = ?", processFile, name).Error
	if err != nil {
		pmond.Log.Infof("could not find process in database: %s or %s %-v", processFile, name, err)
		return nil, err
	}
	return &found, nil
}

func (pr *ProcessRepo) FindByStatus(status model.ProcessStatus) ([]model.Process, error) {
	var all []model.Process
	err := pr.getDb().Find(&all, "status = ?", status).Error
	if err != nil {
		pmond.Log.Infof("pmon3 can find processes with status %s: %v", status, err)
	}
	return all, err
}

func (pr *ProcessRepo) FindForMonitor() ([]model.Process, error) {
	var all []model.Process
	err := pr.getDb().Find(&all, "status in (?, ?, ?, ?, ?)",
		model.StatusRunning,
		model.StatusFailed,
		model.StatusQueued,
		model.StatusClosed,
		model.StatusBackoff,
	).Error
	return all, err
}

func (pr *ProcessRepo) FindAll() ([]model.Process, error) {
	var all []model.Process
	err := pr.getDb().Find(&all).Error
	if err != nil {
		pmond.Log.Infof("cant find processes: %v", err)
	}
	return all, err
}

func (pr *ProcessRepo) FindAllOrdered(orderBy string) ([]model.Process, error) {
	var all []model.Process
	err := pr.getDb().Order(fmt.Sprintf("%s asc", orderBy)).Find(&all).Error
	if err != nil {
		pmond.Log.Infof("cant find processes: %v", err)
	}
	return all, err
}

func (pr *ProcessRepo) SaveGroups() error {
	currentProcess, err := pr.FindById(pr.cur.ID)
	if err != nil {
		return err
	}

	existingLen := len(currentProcess.Groups)
	desiredLen := len(pr.cur.Groups)

	if existingLen == 0 && desiredLen == 0 {
		//there are no groups to add so we return early
		pmond.Log.Infof("no groups to add")
	} else if existingLen == 0 && desiredLen > 0 {
		pmond.Log.Infof("adding groups %-v %-v", currentProcess.Groups, pr.cur.Groups)
		err = pr.db.Model(&pr.cur).Association("Groups").Append(pr.cur.Groups)
	} else if (existingLen > 0 && desiredLen == 0) || (existingLen != desiredLen) {
		if desiredLen > 0 {
			pmond.Log.Infof("replacing groups %-v %-v", currentProcess.Groups, pr.cur.Groups)
			err = pr.db.Model(&pr.cur).Association("Groups").Replace(pr.cur.Groups)
		} else {
			pmond.Log.Infof("purging groups %-v %-v", currentProcess.Groups, pr.cur.Groups)
			err = pr.db.Model(&pr.cur).Association("Groups").Clear()
		}
	} else if existingLen == desiredLen {
		//are they the same set?
		existingHashSet := currentProcess.GetGroupHashSet()
		desiredHashSet := pr.cur.GetGroupHashSet()

		if existingHashSet.Equal(desiredHashSet) {
			pmond.Log.Infof("groups hash sets match")
		} else {
			pmond.Log.Infof("groups hash sets dont match, replacing groups")
			//sets are not equivalent, so we need to replace the groups
			err = pr.db.Model(&pr.cur).Association("Groups").Replace(pr.cur.Groups)
		}
	}

	return err
}

func (pr *ProcessRepo) FindAndSave() (string, error) {

	found, err := pr.FindByFileAndName(pr.cur.ProcessFile, pr.cur.Name)
	if err == nil && found.ID > 0 { // process already exists
		pr.cur.ID = found.ID
	}

	err = pr.Save()
	if err != nil {
		return "", fmt.Errorf("could not save process: err: %w", err)
	}

	err = pr.SaveGroups()
	if err != nil {
		return "", fmt.Errorf("could not save process groups: err: %w", err)
	}

	output, err := json.Marshal(pr.cur.RenderTable())
	return string(output), err
}
