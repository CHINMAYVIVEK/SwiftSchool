package mapper

import (
	"database/sql"
	"fmt"
	"swiftschool/domain"
	"swiftschool/helper"
	"swiftschool/internal/db"

	"github.com/google/uuid"
)

// =========================================================
// FEE HEAD MAPPERS
// =========================================================

func MapFeeHeadDomainToParams(fh domain.FeeHead) db.CreateFeeHeadParams {
	return db.CreateFeeHeadParams{
		InstituteID:       fh.InstituteID,
		Name:              fh.Name,
		IsRefundable:      sql.NullBool{Bool: fh.IsRefundable, Valid: true},
		LinkedGlAccountID: helper.ToNullUUID(fh.LinkedGLAccountID),
		CreatedBy:         helper.ToNullUUID(helper.DerefUUID(fh.CreatedBy)),
	}
}

func MapFeeHeadRowToDomain(row db.FinanceFeeHead) domain.FeeHead {
	return domain.FeeHead{
		TenantUUIDModel: domain.TenantUUIDModel{
			BaseUUIDModel: domain.BaseUUIDModel{
				ID:        row.ID,
				CreatedAt: helper.NullTimeToValue(row.CreatedAt),
				UpdatedAt: helper.NullTimeToValue(row.UpdatedAt),
				DeletedAt: helper.NullTimeToPtr(row.DeletedAt),
				CreatedBy: helper.NullUUIDToPtr(row.CreatedBy),
			},
			InstituteID: row.InstituteID,
		},
		Name:              row.Name,
		IsRefundable:      row.IsRefundable.Bool,
		LinkedGLAccountID: helper.NullUUIDToValue(row.LinkedGlAccountID),
	}
}

// =========================================================
// FEE STRUCTURE MAPPERS
// =========================================================

func MapFeeStructureDomainToParams(fs domain.FeeStructure) db.CreateFeeStructureParams {
	return db.CreateFeeStructureParams{
		InstituteID:       fs.InstituteID,
		AcademicSessionID: fs.AcademicSessionID,
		ClassID:           helper.ToNullUUID(helper.DerefUUID(fs.ClassID)),
		FeeHeadID:         fs.FeeHeadID,
		Amount:            fmt.Sprintf("%.2f", fs.Amount),
		Frequency:         helper.ToNullString(string(fs.Frequency)),
		CreatedBy:         helper.ToNullUUID(helper.DerefUUID(fs.CreatedBy)),
	}
}

func MapFeeStructureRowToDomain(row db.FinanceFeeStructure) domain.FeeStructure {
	var amount float64
	fmt.Sscanf(row.Amount, "%f", &amount)

	return domain.FeeStructure{
		TenantUUIDModel: domain.TenantUUIDModel{
			BaseUUIDModel: domain.BaseUUIDModel{
				ID:        row.ID,
				CreatedAt: helper.NullTimeToValue(row.CreatedAt),
				UpdatedAt: helper.NullTimeToValue(row.UpdatedAt),
				CreatedBy: helper.NullUUIDToPtr(row.CreatedBy),
			},
			InstituteID: row.InstituteID,
		},
		AcademicSessionID: row.AcademicSessionID,
		ClassID:           helper.NullUUIDToPtr(row.ClassID),
		FeeHeadID:         row.FeeHeadID,
		Amount:            amount,
		Frequency:         domain.FeeFrequency(row.Frequency.String),
	}
}

// =========================================================
// FINE RULE MAPPERS
// =========================================================

func MapFineRuleDomainToParams(fr domain.FineRule) db.CreateFineRuleParams {
	return db.CreateFineRuleParams{
		InstituteID: fr.InstituteID,
		Name:        helper.ToNullString(fr.Name),
		GraceDays:   sql.NullInt32{Int32: int32(fr.GraceDays), Valid: true},
		FineType:    helper.ToNullString(string(fr.FineType)),
		FineAmount:  helper.ToNullString(fmt.Sprintf("%.2f", fr.FineAmount)),
		IsActive:    sql.NullBool{Bool: fr.IsActive, Valid: true},
		CreatedBy:   helper.ToNullUUID(helper.DerefUUID(fr.CreatedBy)),
	}
}

func MapFineRuleRowToDomain(row db.FinanceFineRule) domain.FineRule {
	var amount float64
	if row.FineAmount.Valid {
		fmt.Sscanf(row.FineAmount.String, "%f", &amount)
	}

	return domain.FineRule{
		TenantUUIDModel: domain.TenantUUIDModel{
			BaseUUIDModel: domain.BaseUUIDModel{
				ID:        row.ID,
				CreatedAt: helper.NullTimeToValue(row.CreatedAt),
				CreatedBy: helper.NullUUIDToPtr(row.CreatedBy),
			},
			InstituteID: row.InstituteID,
		},
		Name:       row.Name.String,
		GraceDays:  int(row.GraceDays.Int32),
		FineType:   domain.FineType(row.FineType.String),
		FineAmount: amount,
		IsActive:   row.IsActive.Bool,
	}
}

// =========================================================
// INVOICE MAPPERS
// =========================================================

func MapInvoiceDomainToParams(inv domain.Invoice) db.CreateInvoiceParams {
	return db.CreateInvoiceParams{
		InstituteID:       inv.InstituteID,
		InvoiceNo:         inv.InvoiceNo,
		StudentID:         inv.StudentID,
		AcademicSessionID: helper.ToNullUUID(uuid.Nil), // Domain doesn't have this field
		TotalAmount:       fmt.Sprintf("%.2f", inv.TotalAmount),
		DueDate:           helper.ToNullTime(helper.TimeOrZero(inv.DueDate)),
		CreatedBy:         helper.ToNullUUID(helper.DerefUUID(inv.CreatedBy)),
	}
}

func MapInvoiceRowToDomain(row db.FinanceInvoice) domain.Invoice {
	var totalAmount, paidAmount float64
	fmt.Sscanf(row.TotalAmount, "%f", &totalAmount)
	if row.PaidAmount.Valid {
		fmt.Sscanf(row.PaidAmount.String, "%f", &paidAmount)
	}

	return domain.Invoice{
		TenantUUIDModel: domain.TenantUUIDModel{
			BaseUUIDModel: domain.BaseUUIDModel{
				ID:        row.ID,
				CreatedAt: helper.NullTimeToValue(row.CreatedAt),
				UpdatedAt: helper.NullTimeToValue(row.UpdatedAt),
				CreatedBy: helper.NullUUIDToPtr(row.CreatedBy),
			},
			InstituteID: row.InstituteID,
		},
		InvoiceNo:      row.InvoiceNo,
		StudentID:      row.StudentID,
		TotalAmount:    totalAmount,
		DiscountAmount: 0, // Not in SQLC row
		FineAmount:     0, // Not in SQLC row
		PaidAmount:     paidAmount,
		Status:         row.Status.String,
		DueDate:        helper.NullTimeToPtr(row.DueDate),
	}
}

// =========================================================
// INVOICE ITEM MAPPERS
// =========================================================

func MapInvoiceItemDomainToParams(item domain.InvoiceItem) db.CreateInvoiceItemParams {
	return db.CreateInvoiceItemParams{
		InstituteID:     item.InstituteID,
		InvoiceID:       item.InvoiceID,
		FeeHeadID:       helper.ToNullUUID(helper.DerefUUID(item.FeeHeadID)),
		Amount:          fmt.Sprintf("%.2f", item.Amount),
		ConcessionID:    helper.ToNullUUID(helper.DerefUUID(item.ConcessionID)),
		DiscountApplied: helper.ToNullString(fmt.Sprintf("%.2f", item.DiscountApplied)),
		Description:     helper.ToNullString(helper.StrOrEmpty(item.Description)),
		CreatedBy:       helper.ToNullUUID(helper.DerefUUID(item.CreatedBy)),
	}
}

func MapInvoiceItemRowToDomain(row db.FinanceInvoiceItem) domain.InvoiceItem {
	var amount, discount float64
	fmt.Sscanf(row.Amount, "%f", &amount)
	if row.DiscountApplied.Valid {
		fmt.Sscanf(row.DiscountApplied.String, "%f", &discount)
	}

	return domain.InvoiceItem{
		TenantUUIDModel: domain.TenantUUIDModel{
			BaseUUIDModel: domain.BaseUUIDModel{
				ID:        row.ID,
				CreatedAt: helper.NullTimeToValue(row.CreatedAt),
				CreatedBy: helper.NullUUIDToPtr(row.CreatedBy),
			},
			InstituteID: row.InstituteID,
		},
		InvoiceID:       row.InvoiceID,
		FeeHeadID:       helper.NullUUIDToPtr(row.FeeHeadID),
		Amount:          amount,
		ConcessionID:    helper.NullUUIDToPtr(row.ConcessionID),
		DiscountApplied: discount,
		Description:     helper.NullStringToPtr(row.Description),
	}
}

// =========================================================
// TRANSACTION MAPPERS
// =========================================================

func MapTransactionDomainToParams(txn domain.Transaction) db.CreateTransactionParams {
	return db.CreateTransactionParams{
		InstituteID:      txn.InstituteID,
		InvoiceID:        helper.ToNullUUID(helper.DerefUUID(txn.InvoiceID)),
		StudentID:        helper.ToNullUUID(helper.DerefUUID(txn.StudentID)),
		TransactionRefNo: helper.ToNullString(helper.StrOrEmpty(txn.TransactionRefNo)),
		PaymentMode:      helper.ToNullString(string(txn.PaymentMode)),
		Amount:           fmt.Sprintf("%.2f", txn.Amount),
		ChequeNo:         helper.ToNullString(helper.StrOrEmpty(txn.ChequeNo)),
		ChequeDate:       helper.ToNullTime(helper.TimeOrZero(txn.ChequeDate)),
		BankName:         helper.ToNullString(helper.StrOrEmpty(txn.BankName)),
		CollectedBy:      helper.ToNullUUID(helper.DerefUUID(txn.CollectedBy)),
		CreatedBy:        helper.ToNullUUID(helper.DerefUUID(txn.CreatedBy)),
	}
}

func MapTransactionRowToDomain(row db.FinanceTransaction) domain.Transaction {
	var amount float64
	fmt.Sscanf(row.Amount, "%f", &amount)

	return domain.Transaction{
		TenantUUIDModel: domain.TenantUUIDModel{
			BaseUUIDModel: domain.BaseUUIDModel{
				ID:        row.ID,
				CreatedAt: helper.NullTimeToValue(row.CreatedAt),
				CreatedBy: helper.NullUUIDToPtr(row.CreatedBy),
			},
			InstituteID: row.InstituteID,
		},
		InvoiceID:        helper.NullUUIDToPtr(row.InvoiceID),
		StudentID:        helper.NullUUIDToPtr(row.StudentID),
		TransactionRefNo: helper.NullStringToPtr(row.TransactionRefNo),
		PaymentMode:      domain.PaymentMode(row.PaymentMode.String),
		Amount:           amount,
		ChequeNo:         helper.NullStringToPtr(row.ChequeNo),
		ChequeDate:       helper.NullTimeToPtr(row.ChequeDate),
		BankName:         helper.NullStringToPtr(row.BankName),
		Status:           row.Status.String,
		CollectedBy:      helper.NullUUIDToPtr(row.CollectedBy),
	}
}

// =========================================================
// ACCOUNT (GL) MAPPERS
// =========================================================

func MapAccountDomainToParams(acc domain.Account) db.CreateAccountParams {
	return db.CreateAccountParams{
		InstituteID:     acc.InstituteID,
		Name:            acc.Name,
		Code:            acc.Code,
		ParentAccountID: helper.ToNullUUID(helper.DerefUUID(acc.ParentAccountID)),
		Type:            string(acc.Type),
		IsSystem:        sql.NullBool{Bool: acc.IsSystem, Valid: true},
		CreatedBy:       helper.ToNullUUID(helper.DerefUUID(acc.CreatedBy)),
	}
}

func MapAccountRowToDomain(row db.FinanceAccount) domain.Account {
	return domain.Account{
		TenantUUIDModel: domain.TenantUUIDModel{
			BaseUUIDModel: domain.BaseUUIDModel{
				ID:        row.ID,
				CreatedAt: helper.NullTimeToValue(row.CreatedAt),
				CreatedBy: helper.NullUUIDToPtr(row.CreatedBy),
			},
			InstituteID: row.InstituteID,
		},
		Name:            row.Name,
		Code:            row.Code,
		ParentAccountID: helper.NullUUIDToPtr(row.ParentAccountID),
		Type:            domain.AccountType(row.Type),
		IsSystem:        row.IsSystem.Bool,
	}
}

// =========================================================
// JOURNAL ENTRY MAPPERS
// =========================================================

func MapJournalEntryDomainToParams(je domain.JournalEntry) db.CreateJournalEntryParams {
	return db.CreateJournalEntryParams{
		InstituteID:     je.InstituteID,
		ReferenceNo:     helper.ToNullString(je.ReferenceNo),
		TransactionDate: helper.ToNullTime(je.TransactionDate),
		Description:     helper.ToNullString(helper.StrOrEmpty(je.Description)),
		IsPosted:        sql.NullBool{Bool: je.IsPosted, Valid: true},
		CreatedBy:       helper.ToNullUUID(helper.DerefUUID(je.CreatedBy)),
	}
}

func MapJournalEntryRowToDomain(row db.FinanceJournalEntry) domain.JournalEntry {
	return domain.JournalEntry{
		TenantUUIDModel: domain.TenantUUIDModel{
			BaseUUIDModel: domain.BaseUUIDModel{
				ID:        row.ID,
				CreatedAt: helper.NullTimeToValue(row.CreatedAt),
				CreatedBy: helper.NullUUIDToPtr(row.CreatedBy),
			},
			InstituteID: row.InstituteID,
		},
		ReferenceNo:     helper.NullStringToValue(row.ReferenceNo),
		TransactionDate: helper.NullTimeToValue(row.TransactionDate),
		Description:     helper.NullStringToPtr(row.Description),
		IsPosted:        row.IsPosted.Bool,
		PostedAt:        helper.NullTimeToPtr(row.PostedAt),
	}
}

// =========================================================
// JOURNAL ITEM MAPPERS
// =========================================================

func MapJournalItemDomainToParams(ji domain.JournalItem) db.CreateJournalItemParams {
	return db.CreateJournalItemParams{
		InstituteID:    ji.InstituteID,
		JournalEntryID: ji.JournalEntryID,
		AccountID:      ji.AccountID,
		Debit:          helper.ToNullString(fmt.Sprintf("%.2f", ji.Debit)),
		Credit:         helper.ToNullString(fmt.Sprintf("%.2f", ji.Credit)),
		Description:    helper.ToNullString(helper.StrOrEmpty(ji.Description)),
		CreatedBy:      helper.ToNullUUID(helper.DerefUUID(ji.CreatedBy)),
	}
}

func MapJournalItemRowToDomain(row db.FinanceJournalItem) domain.JournalItem {
	var debit, credit float64
	if row.Debit.Valid {
		fmt.Sscanf(row.Debit.String, "%f", &debit)
	}
	if row.Credit.Valid {
		fmt.Sscanf(row.Credit.String, "%f", &credit)
	}

	return domain.JournalItem{
		TenantUUIDModel: domain.TenantUUIDModel{
			BaseUUIDModel: domain.BaseUUIDModel{
				ID:        row.ID,
				CreatedAt: helper.NullTimeToValue(row.CreatedAt),
				CreatedBy: helper.NullUUIDToPtr(row.CreatedBy),
			},
			InstituteID: row.InstituteID,
		},
		JournalEntryID: row.JournalEntryID,
		AccountID:      row.AccountID,
		Debit:          debit,
		Credit:         credit,
		Description:    helper.NullStringToPtr(row.Description),
	}
}

// =========================================================
// VENDOR MAPPERS
// =========================================================

func MapVendorDomainToParams(v domain.Vendor) db.CreateVendorParams {
	return db.CreateVendorParams{
		InstituteID: v.InstituteID,
		Name:        v.Name,
		ContactName: helper.ToNullString(helper.StrOrEmpty(v.ContactName)),
		Phone:       helper.ToNullString(helper.StrOrEmpty(v.Phone)),
		Email:       helper.ToNullString(helper.StrOrEmpty(v.Email)),
		Address:     helper.ToNullString(helper.StrOrEmpty(v.Address)),
		GstNumber:   helper.ToNullString(helper.StrOrEmpty(v.GSTNumber)),
		CreatedBy:   helper.ToNullUUID(helper.DerefUUID(v.CreatedBy)),
	}
}

func MapVendorRowToDomain(row db.FinanceVendor) domain.Vendor {
	return domain.Vendor{
		TenantUUIDModel: domain.TenantUUIDModel{
			BaseUUIDModel: domain.BaseUUIDModel{
				ID:        row.ID,
				CreatedAt: helper.NullTimeToValue(row.CreatedAt),
				CreatedBy: helper.NullUUIDToPtr(row.CreatedBy),
			},
			InstituteID: row.InstituteID,
		},
		Name:        row.Name,
		ContactName: helper.NullStringToPtr(row.ContactName),
		Phone:       helper.NullStringToPtr(row.Phone),
		Email:       helper.NullStringToPtr(row.Email),
		Address:     helper.NullStringToPtr(row.Address),
		GSTNumber:   helper.NullStringToPtr(row.GstNumber),
	}
}

// =========================================================
// PURCHASE ORDER MAPPERS
// =========================================================

func MapPurchaseOrderDomainToParams(po domain.PurchaseOrder) db.CreatePurchaseOrderParams {
	return db.CreatePurchaseOrderParams{
		InstituteID: po.InstituteID,
		VendorID:    po.VendorID,
		OrderDate:   helper.ToNullTime(po.OrderDate),
		CreatedBy:   helper.ToNullUUID(helper.DerefUUID(po.CreatedBy)),
	}
}

func MapPurchaseOrderRowToDomain(row db.FinancePurchaseOrder) domain.PurchaseOrder {
	var totalAmount float64
	if row.TotalAmount.Valid {
		fmt.Sscanf(row.TotalAmount.String, "%f", &totalAmount)
	}

	return domain.PurchaseOrder{
		TenantUUIDModel: domain.TenantUUIDModel{
			BaseUUIDModel: domain.BaseUUIDModel{
				ID:        row.ID,
				CreatedAt: helper.NullTimeToValue(row.CreatedAt),
				CreatedBy: helper.NullUUIDToPtr(row.CreatedBy),
			},
			InstituteID: row.InstituteID,
		},
		VendorID:    row.VendorID,
		OrderDate:   helper.NullTimeToValue(row.OrderDate),
		Status:      domain.PurchaseStatus(row.Status.String),
		TotalAmount: totalAmount,
		ReferenceNo: helper.NullStringToPtr(row.ReferenceNo),
	}
}

// =========================================================
// PURCHASE ITEM MAPPERS
// =========================================================

func MapPurchaseItemDomainToParams(pi domain.PurchaseItem) db.AddPurchaseItemParams {
	return db.AddPurchaseItemParams{
		InstituteID:     helper.ToNullUUID(pi.InstituteID),
		PurchaseOrderID: helper.ToNullUUID(pi.PurchaseOrderID),
		ItemID:          helper.ToNullUUID(pi.ItemID),
		Quantity:        sql.NullInt32{Int32: int32(pi.Quantity), Valid: true},
		UnitPrice:       helper.ToNullString(fmt.Sprintf("%.2f", pi.UnitPrice)),
		TaxID:           helper.ToNullUUID(uuid.Nil), // Placeholder
		TaxAmount:       helper.ToNullString("0"),
		TotalAmount:     helper.ToNullString(fmt.Sprintf("%.2f", float64(pi.Quantity)*pi.UnitPrice)),
	}
}

func MapPurchaseItemRowToDomain(row db.FinancePurchaseOrderItem) domain.PurchaseItem {
	var unitPrice float64
	if row.UnitPrice.Valid {
		fmt.Sscanf(row.UnitPrice.String, "%f", &unitPrice)
	}

	return domain.PurchaseItem{
		TenantUUIDModel: domain.TenantUUIDModel{
			BaseUUIDModel: domain.BaseUUIDModel{
				ID:        row.ID,
				CreatedAt: helper.NullTimeToValue(row.CreatedAt),
			},
			InstituteID: helper.NullUUIDToValue(row.InstituteID),
		},
		PurchaseOrderID: helper.NullUUIDToValue(row.PurchaseOrderID),
		ItemID:          helper.NullUUIDToValue(row.ItemID),
		Quantity:        int(row.Quantity.Int32),
		UnitPrice:       unitPrice,
	}
}
