package controllers

import "tech-challenge/src/models"

type ServiceController struct {
	repo models.Repository
}

func NewServiceController(repo models.Repository) *ServiceController {
	return &ServiceController{repo: repo}
}

func (sc *ServiceController) getEmails() ([]models.Email, error) {
	return sc.repo.getEmails()
}

func (sc *ServiceController) getEmailById(id int) ([]models.Email, error) {
	return sc.repo.getEmailById(id)
}

func (sc *ServiceController) createEmail(email *models.Email) ([]models.Email, error) {
	return sc.repo.createEmail(user)
}

func (sc *ServiceController) updateEmail(email *models.Email) ([]models.Email, error) {
	return sc.repo.updateEmail(user)
}

func (sc *ServiceController) deleteEmail(id int) ([]models.Email, error) {
	return sc.repo.deleteEmail(id)
}
