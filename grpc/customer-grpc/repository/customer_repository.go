package customer_repo

import (
	"context"
	"mock-golang/database"
	customer_model "mock-golang/grpc/customer-grpc/model"
	customer_request "mock-golang/grpc/customer-grpc/request"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

//Embeded struct

type CustomerRepository interface {
	FindById(ctx context.Context, id uuid.UUID) (*customer_model.Customer, error)
	CreateCustomer(ctx context.Context, model *customer_model.Customer) (*customer_model.Customer, error)
	UpdateCustomer(ctx context.Context, model *customer_model.Customer) (*customer_model.Customer, error)
	SearchCustomer(ctx context.Context, req *customer_request.SearchCustomerRequest) ([]*customer_model.Customer, error)
}

type dbmanager struct {
	*gorm.DB
}

func NewDBManager() (CustomerRepository, error) {
	db, err := database.NewGormDB()
	if err != nil {
		return nil, err
	}

	db = db.Debug()

	err = db.AutoMigrate(
		&customer_model.Customer{},
	)

	if err != nil {
		return nil, err
	}

	return &dbmanager{db}, nil
}

func (m *dbmanager) FindById(ctx context.Context, id uuid.UUID) (*customer_model.Customer, error) {
	res := customer_model.Customer{}
	if err := m.Where(&customer_model.Customer{Id: id}).First(&res).Error; err != nil {
		return nil, err
	}

	return &res, nil
}

func (m *dbmanager) CreateCustomer(ctx context.Context, model *customer_model.Customer) (*customer_model.Customer, error) {
	if err := m.Create(model).Error; err != nil {
		return nil, err
	}

	return model, nil
}

func (m *dbmanager) UpdateCustomer(ctx context.Context, model *customer_model.Customer) (*customer_model.Customer, error) {
	if err := m.Where(&customer_model.Customer{Id: model.Id}).Updates(model).Error; err != nil {
		return nil, err
	}

	return model, nil
}

func (m *dbmanager) SearchCustomer(ctx context.Context, req *customer_request.SearchCustomerRequest) ([]*customer_model.Customer, error) {
	customers := []*customer_model.Customer{}

	sbWhere := " 1=1 "
	params := []interface{}{}
	if len(strings.TrimSpace(req.Id)) > 0 {
		sbWhere += " AND Id = ? "
		params = append(params, req.Id)
	}
	if len(strings.TrimSpace(req.Name)) > 0 {
		sbWhere += " AND Name = ? "
		params = append(params, req.Name)
	}
	if req.Role >= 0 {
		sbWhere += " AND Role = ? "
		params = append(params, req.Role)
	}
	if len(strings.TrimSpace(req.Email)) > 0 {
		sbWhere += " AND Email = ? "
		params = append(params, req.Email)
	}
	if len(strings.TrimSpace(req.PhoneNumber)) > 0 {
		sbWhere += " AND phone_number = ? "
		params = append(params, req.PhoneNumber)
	}
	if len(strings.TrimSpace(req.IdentityCard)) > 0 {
		sbWhere += " AND identity_card = ? "
		params = append(params, req.IdentityCard)
	}
	if len(strings.TrimSpace(req.MembershipCard)) > 0 {
		sbWhere += " AND membership_card = ? "
		params = append(params, req.MembershipCard)
	}
	if req.Status >= 0 {
		sbWhere += " AND Status = ? "
		params = append(params, req.Status)
	}

	if err := m.Where(sbWhere, params...).Find(&customers).Error; err != nil {
		return nil, err
	}

	return customers, nil
}
