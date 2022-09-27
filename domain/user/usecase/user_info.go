package usecase

import (
	"context"
	"lucy/cashier/domain"
	"lucy/cashier/dto"
	"net/http"
	"time"
)

func (u *userUsecase) UserBranchInfo(c context.Context, id string, withTrashed bool) (*dto.BranchResponse, int, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	user, code, err := u.userRepo.FindUserBy(ctx, "uuid", id, withTrashed)
	if err != nil {
		return nil, code, err
	}

	company, _, _ := u.companyRepo.FindCompanyByBranchUUID(ctx, user.BranchUUID, true)
	branchDiscounts, _, _ := u.branchDiscountRepo.FindBranchDiscounts(ctx, user.BranchUUID)

	var resp dto.BranchResponse
	userDomainToDTO_UserBranchInfo(&resp, user, company, branchDiscounts)

	return &resp, http.StatusOK, nil
}

func userDomainToDTO_UserBranchInfo(resp *dto.BranchResponse, u *domain.User, c *domain.Company, bd []domain.BranchDiscount) {
	// branch
	for _, branch := range c.Branches {
		if branch.UUID != u.BranchUUID {
			continue
		}

		resp.UUID = branch.UUID
		resp.UniqueIndentifier = branch.UniqueIndentifier
		resp.Name = branch.Name
		resp.Timezone = branch.Timezone
		resp.Public = branch.Public
		resp.Preferences = branch.Preferences
		resp.Phone = branch.Phone
		resp.CreatedAt = time.UnixMicro(branch.CreatedAt).UTC()

		// address
		resp.Address.Province = branch.Address.Province
		resp.Address.City = branch.Address.City
		resp.Address.Street = branch.Address.Street
		resp.Address.PostalCode = branch.Address.PostalCode

		// company
		resp.Company.UUID = c.UUID
		resp.Company.Name = c.Name
		resp.Company.CreatedAt = time.UnixMicro(c.CreatedAt).UTC()

		// social media
		resp.SocialMedia.Facebook = branch.SocialMedia.Facebook
		resp.SocialMedia.Instagram = branch.SocialMedia.Instagram
		resp.SocialMedia.Twitter = branch.SocialMedia.Twitter
		resp.SocialMedia.Tiktok = branch.SocialMedia.Tiktok
		resp.SocialMedia.GoogleMaps = branch.SocialMedia.GoogleMaps
		if branch.SocialMedia.Whatsapp != nil {
			var whatsapp dto.WhatsappResponse
			whatsapp.CountryCode = *&branch.SocialMedia.Whatsapp.CountryCode
			whatsapp.Number = *&branch.SocialMedia.Whatsapp.Number

			resp.SocialMedia.Whatsapp = &whatsapp
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
		resp.Discounts = discounts
		if len(discounts) == 0 {
			resp.Discounts = make([]dto.BranchDiscountResponse, 0)
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
		resp.Taxes = respTaxes
		if len(respTaxes) == 0 {
			resp.Taxes = make([]dto.TaxResponse, 0)
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
		resp.Tips = respTips
		if len(respTips) == 0 {
			resp.Tips = make([]dto.TipResponse, 0)
		}

		// payment methods
		var respPaymentMethods []dto.PaymentMethodResponse
		for _, paymentMethod := range branch.PaymentMethods {
			var paymentMethodResp dto.PaymentMethodResponse
			paymentMethodResp.UUID = paymentMethod.UUID
			paymentMethodResp.EntryUUID = paymentMethod.EntryUUID
			paymentMethodResp.Group = paymentMethod.Group
			paymentMethodResp.Name = paymentMethod.Name
			paymentMethodResp.Fee = paymentMethod.Fee
			paymentMethodResp.Description = paymentMethod.Description
			paymentMethodResp.Disabled = paymentMethod.Disabled
			paymentMethodResp.CreatedAt = time.UnixMicro(paymentMethod.CreatedAt).UTC()
			if paymentMethod.UpdatedAt != nil {
				paymentMethodRespUpdatedAt := time.UnixMicro(*paymentMethod.UpdatedAt).UTC()
				paymentMethodResp.UpdatedAt = &paymentMethodRespUpdatedAt
			}
			if paymentMethod.DeletedAt != nil {
				paymentMethodRespDeletedAt := time.UnixMicro(*paymentMethod.DeletedAt).UTC()
				paymentMethodResp.DeletedAt = &paymentMethodRespDeletedAt
			}

			respPaymentMethods = append(respPaymentMethods, paymentMethodResp)
		}
		resp.PaymentMethods = respPaymentMethods
		if len(respPaymentMethods) == 0 {
			resp.PaymentMethods = make([]dto.PaymentMethodResponse, 0)
		}

		//

		// fee preferences
		resp.FeePreference.Service.Nominal = branch.FeePreference.Service.Nominal
		resp.FeePreference.Service.Percentage = branch.FeePreference.Service.Percentage

		resp.FeePreference.Queue.Nominal = branch.FeePreference.Queue.Nominal
		resp.FeePreference.Queue.Percentage = branch.FeePreference.Queue.Percentage

		resp.FeePreference.Gojek.Nominal = branch.FeePreference.Gojek.Nominal
		resp.FeePreference.Gojek.Percentage = branch.FeePreference.Gojek.Percentage

		resp.FeePreference.Grab.Nominal = branch.FeePreference.Grab.Nominal
		resp.FeePreference.Grab.Percentage = branch.FeePreference.Grab.Percentage

		resp.FeePreference.Shopee.Nominal = branch.FeePreference.Shopee.Nominal
		resp.FeePreference.Shopee.Percentage = branch.FeePreference.Shopee.Percentage

		resp.FeePreference.Maxim.Nominal = branch.FeePreference.Maxim.Nominal
		resp.FeePreference.Maxim.Percentage = branch.FeePreference.Maxim.Percentage

		resp.FeePreference.Private.Nominal = branch.FeePreference.Private.Nominal
		resp.FeePreference.Private.Percentage = branch.FeePreference.Private.Percentage

		// employees
		resp.Employees = make([]dto.UserResponse, 0)
	}
}
