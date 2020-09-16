package backup

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/sirupsen/logrus"
	"github.com/ticketmaster/lbapi/config"
	"github.com/ticketmaster/lbapi/shared"
)

var log = logrus.New().WithField("route", "backup")

type Backup struct {
	User       string
	FullName   string
	Email      string
	Password   string
	Remote     string
	RepoName   string
	Enabled    bool
	Repository *git.Repository
}

func New() *Backup {
	return &Backup{
		User:     config.GlobalConfig.Backup.User,
		FullName: config.GlobalConfig.Backup.FullName,
		Email:    config.GlobalConfig.Backup.Email,
		Password: config.GlobalConfig.Backup.Password,
		Remote:   config.GlobalConfig.Backup.Remote,
		Enabled:  config.GlobalConfig.Backup.Enable,
		RepoName: config.GlobalConfig.Backup.RepoName,
	}
}

func (o *Backup) Commit(in interface{}) (err error) {
	if o.Enabled == false {
		return
	}
	var dbRecordCollection DbRecordCollection
	err = shared.MarshalInterface(in, &dbRecordCollection)
	if err != nil {
		return err
	}
	////////////////////////////////////////////////////////////////////////////
	var w *git.Worktree
	if o.Repository == nil {
		o.Repository, err = git.PlainClone(o.RepoName, false, &git.CloneOptions{URL: o.Remote})
		if err != nil {
			msg := err.Error()
			if strings.Contains(msg, "repository already exists") {
				log.Warn(err)
				o.Repository, err = git.PlainOpen(o.RepoName)
				if err != nil {
					return err
				}
				w, err = o.Repository.Worktree()
				if err != nil {
					return err
				}
				err = w.Pull(&git.PullOptions{RemoteName: "origin"})
				if err != nil {
					msg := err.Error()
					if strings.Contains(msg, "already up-to-date") {
						log.Warn(err)
					} else {
						return err
					}
				}
			} else {
				return err
			}
		}
		////////////////////////////////////////////////////////////////////////
		o.Repository, err = git.PlainOpen(o.RepoName)
		if err != nil {
			return err
		}
		////////////////////////////////////////////////////////////////////////
		w, err = o.Repository.Worktree()
		if err != nil {
			return err
		}

	}
	////////////////////////////////////////////////////////////////////////////
	for _, v := range dbRecordCollection.DbRecords {
		data, err := json.Marshal(v.Data)
		if err != nil {
			return err
		}
		////////////////////////////////////////////////////////////////////////
		f := filepath.Join(o.RepoName, v.ID+".json")
		err = ioutil.WriteFile(f, data, 0644)
		if err != nil {
			return err
		}
		////////////////////////////////////////////////////////////////////////
		_, err = w.Add(v.ID + ".json")
		if err != nil {
			log.Warn(err)
		}
	}
	////////////////////////////////////////////////////////////////////////////
	commit, err := w.Commit("update", &git.CommitOptions{
		All: true,
		Author: &object.Signature{
			Name:  o.FullName,
			Email: o.Email,
			When:  time.Now(),
		},
	})
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////////////////////////
	_, err = o.Repository.CommitObject(commit)
	if err != nil {
		return
	}
	err = o.Repository.Push(&git.PushOptions{
		Auth: &http.BasicAuth{
			Username: o.User,
			Password: o.Password,
		},
	})
	if err != nil {
		return
	}
	return
}
