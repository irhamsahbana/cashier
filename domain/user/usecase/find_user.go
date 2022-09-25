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
	branchDiscounts, _, _ := u.branchDiscountRepo.FindBranchDiscounts(ctx, user.BranchUUID)

	var resp dto.UserResponse
	userDomainToDTO_FindUser(&resp, user, role, company, branchDiscounts)

	return &resp, http.StatusOK, nil
}

func userDomainToDTO_FindUser(resp *dto.UserResponse, u *domain.User, r *domain.UserRole, c *domain.Company, bd []domain.BranchDiscount) {
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

		// discount
		var discounts []dto.BranchDiscountResponse
		for _, discount := range bd {
			var discountDTO dto.BranchDiscountResponse

			discountDTO.UUID = discount.UUID
			discountDTO.Name = discount.Name
			discountDTO.Description = discount.Description
			discountDTO.Fixed = discount.Fixed
			discountDTO.Percentage = discount.Percentage
			discountDTO.CreatedAt = time.UnixMicro(discount.CreatedAt).UTC()

			if discount.UpdatedAt != nil {
				bdUpdatedAt := time.UnixMicro(*discount.UpdatedAt).UTC()
				discountDTO.UpdatedAt = &bdUpdatedAt
			}

			if discount.DeletedAt != nil {
				bdDeletedAt := time.UnixMicro(*discount.DeletedAt).UTC()
				discountDTO.DeletedAt = &bdDeletedAt
			}

			discounts = append(discounts, discountDTO)
		}
		resp.Branch.Discounts = discounts
		if len(discounts) == 0 {
			resp.Branch.Discounts = make([]dto.BranchDiscountResponse, 0)
		}

		// taxes
		var respTaxes []dto.TaxResponse
		for _, tax := range branch.Taxes {
			var taxResp dto.TaxResponse

			taxResp.UUID = tax.UUID
			taxResp.Name = tax.Name
			taxResp.Value = tax.Value
			taxResp.Description = tax.Description
			taxResp.CreatedAt = time.UnixMicro(tax.CreatedAt).UTC()
			if tax.UpdatedAt != nil {
				taxRespUpdatedAt := time.UnixMicro(*tax.UpdatedAt).UTC()
				taxResp.UpdatedAt = &taxRespUpdatedAt
			}
			if tax.DeletedAt != nil {
				taxRespDeletedAt := time.UnixMicro(*tax.DeletedAt).UTC()
				taxResp.DeletedAt = &taxRespDeletedAt
			}

			respTaxes = append(respTaxes, taxResp)
		}
		resp.Branch.Taxes = respTaxes
		if len(respTaxes) == 0 {
			resp.Branch.Taxes = make([]dto.TaxResponse, 0)
		}

		// tips
		var respTips []dto.TipResponse
		for _, tip := range branch.Tips {
			var tipResp dto.TipResponse
			tipResp.UUID = tip.UUID
			tipResp.Name = tip.Name
			tipResp.Value = tip.Value
			tipResp.Description = tip.Description
			tipResp.CreatedAt = time.UnixMicro(tip.CreatedAt).UTC()
			if tip.UpdatedAt != nil {
				tipRespUpdatedAt := time.UnixMicro(*tip.UpdatedAt).UTC()
				tipResp.UpdatedAt = &tipRespUpdatedAt
			}
			if tip.DeletedAt != nil {
				tipRespDeletedAt := time.UnixMicro(*tip.DeletedAt).UTC()
				tipResp.DeletedAt = &tipRespDeletedAt
			}

			respTips = append(respTips, tipResp)
		}
		resp.Branch.Tips = respTips
		if len(respTips) == 0 {
			resp.Branch.Tips = make([]dto.TipResponse, 0)
		}

		// fee preferences
		resp.Branch.FeePreference.Service.Nominal = branch.FeePreference.Service.Nominal
		resp.Branch.FeePreference.Service.Percentage = branch.FeePreference.Service.Percentage

		resp.Branch.FeePreference.Queue.Nominal = branch.FeePreference.Queue.Nominal
		resp.Branch.FeePreference.Queue.Percentage = branch.FeePreference.Queue.Percentage

		resp.Branch.FeePreference.Gojek.Nominal = branch.FeePreference.Gojek.Nominal
		resp.Branch.FeePreference.Gojek.Percentage = branch.FeePreference.Gojek.Percentage

		resp.Branch.FeePreference.Grab.Nominal = branch.FeePreference.Grab.Nominal
		resp.Branch.FeePreference.Grab.Percentage = branch.FeePreference.Grab.Percentage

		resp.Branch.FeePreference.Shopee.Nominal = branch.FeePreference.Shopee.Nominal
		resp.Branch.FeePreference.Shopee.Percentage = branch.FeePreference.Shopee.Percentage

		resp.Branch.FeePreference.Maxim.Nominal = branch.FeePreference.Maxim.Nominal
		resp.Branch.FeePreference.Maxim.Percentage = branch.FeePreference.Maxim.Percentage

		resp.Branch.FeePreference.Private.Nominal = branch.FeePreference.Private.Nominal
		resp.Branch.FeePreference.Private.Percentage = branch.FeePreference.Private.Percentage
	}
}
