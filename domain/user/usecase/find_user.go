package usecase

import (
	"context"
	"lucy/cashier/domain"
	"lucy/cashier/dto"
	"net/http"
	"time"
)

func (u *userUsecase) FindUser(c context.Context, id string, withTrashed bool) (*dto.UserResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	user, code, err := u.userRepo.FindUserBy(ctx, "uuid", id, withTrashed)
	if err != nil {
		return nil, code, err
	}

	role, _, _ := u.userRoleRepo.FindUserRole(ctx, user.RoleUUID, true)
	company, _, _ := u.companyRepo.FindCompanyByBranchUUID(ctx, user.BranchUUID, true)

	var resp dto.UserResponse
	userDomainToDTO_FindUser(&resp, user, role, company)

	return &resp, http.StatusOK, nil
}

func userDomainToDTO_FindUser(resp *dto.UserResponse, u *domain.User, r *domain.UserRole, c *domain.Company) {
	resp.UUID = u.UUID
	resp.Name = u.Name
	resp.BranchUUID = u.BranchUUID

	respCreatedAt := time.UnixMicro(u.CreatedAt).UTC()
	resp.CreatedAt = &respCreatedAt

	if u.UpdatedAt != nil {
		respUpdatedAt := time.UnixMicro(*u.UpdatedAt).UTC()
		resp.UpdatedAt = &respUpdatedAt
	}

	if u.DeletedAt != nil {
		respDeletedAt := time.UnixMicro(*u.DeletedAt).UTC()
		resp.DeletedAt = &respDeletedAt
	}

	// role
	if r != nil {
		resp.Role = r.Name
	}

	// branch
	for _, branch := range c.Branches {
		if branch.UUID != u.BranchUUID {
			continue
		}

		// init branch pointer
		var b dto.BranchResponse
		resp.Branch = &b

		resp.Branch.UUID = branch.UUID
		resp.Branch.UniqueIndentifier = branch.UniqueIndentifier
		resp.Branch.Name = branch.Name
		resp.Branch.Timezone = branch.Timezone
		resp.Branch.Public = branch.Public
		resp.Branch.Preferences = branch.Preferences
		resp.Branch.CreatedAt = time.UnixMicro(branch.CreatedAt).UTC()

		// address
		resp.Branch.Address.Province = branch.Address.Province
		resp.Branch.Address.City = branch.Address.City
		resp.Branch.Address.Street = branch.Address.Street
		resp.Branch.Address.PostalCode = branch.Address.PostalCode

		// company
		resp.Branch.Company.UUID = c.UUID
		resp.Branch.Company.Name = c.Name
		resp.Branch.Company.CreatedAt = time.UnixMicro(c.CreatedAt).UTC()

		// social media
		resp.Branch.SocialMedia.Facebook = branch.SocialMedia.Facebook
		resp.Branch.SocialMedia.Instagram = branch.SocialMedia.Instagram
		resp.Branch.SocialMedia.Twitter = branch.SocialMedia.Twitter
		resp.Branch.SocialMedia.Tiktok = branch.SocialMedia.Tiktok
		resp.Branch.SocialMedia.GoogleMaps = branch.SocialMedia.GoogleMaps
		if branch.SocialMedia.Whatsapp != nil {
			var whatsapp dto.WhatsappResponse
			whatsapp.CountryCode = *&branch.SocialMedia.Whatsapp.CountryCode
			whatsapp.Number = *&branch.SocialMedia.Whatsapp.Number

			resp.Branch.SocialMedia.Whatsapp = &whatsapp
		}
	}
}
