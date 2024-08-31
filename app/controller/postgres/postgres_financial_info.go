// PostgreSQL を利用するための関数をまとめたパッケージ
package postgres

import (
	log "app/controller/log"
	model "app/domain/model"

	_ "github.com/lib/pq"
)

/*
財務情報テーブルに INSERT する関数
  - return) financial	財務情報
  - return) err			エラー
*/
func InsertFinancialInfo(financial model.FinancialInfo) (err error) {
	// 財務情報テーブルに INSERT
	_, err = db.Exec("INSERT INTO financial_info (code, disclosed_date, disclosed_time, net_sales, operating_profit, ordinary_profit, profit, earnings_per_share, total_assets, equity, equity_to_asset_ratio, book_value_per_share, cash_flows_from_operating_activities, cash_flows_from_investing_activities, cash_flows_from_financing_activities, cash_and_equivalents, result_dividend_per_share_annual, result_payout_ratio_annual, forecast_dividend_per_share_annual, forecast_payout_ratio_annual, next_year_forecast_dividend_per_share_annual, next_year_forecast_payout_ratio_annual, forecast_net_sales, forecast_operating_profit, forecast_ordinary_profit, forecast_profit, forecast_earnings_per_share, next_year_forecast_net_sales, next_year_forecast_operating_profit, next_year_forecast_ordinary_profit, next_year_forecast_profit, next_year_forecast_earnings_per_share, number_of_issued_and_outstanding_shares_at_the_end_of_fiscal_year_including_treasury_stock) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33)",
		financial.Code,
		financial.DisclosedDate,
		financial.DisclosedTime,
		financial.NetSales,
		financial.OperatingProfit,
		financial.OrdinaryProfit,
		financial.Profit,
		financial.EarningsPerShare,
		financial.TotalAssets,
		financial.Equity,
		financial.EquityToAssetRatio,
		financial.BookValuePerShare,
		financial.CashFlowsFromOperatingActivities,
		financial.CashFlowsFromInvestingActivities,
		financial.CashFlowsFromFinancingActivities,
		financial.CashAndEquivalents,
		financial.ResultDividendPerShareAnnual,
		financial.ResultPayoutRatioAnnual,
		financial.ForecastDividendPerShareAnnual,
		financial.ForecastPayoutRatioAnnual,
		financial.NextYearForecastDividendPerShareAnnual,
		financial.NextYearForecastPayoutRatioAnnual,
		financial.ForecastNetSales,
		financial.ForecastOperatingProfit,
		financial.ForecastOrdinaryProfit,
		financial.ForecastProfit,
		financial.ForecastEarningsPerShare,
		financial.NextYearForecastNetSales,
		financial.NextYearForecastOperatingProfit,
		financial.NextYearForecastOrdinaryProfit,
		financial.NextYearForecastProfit,
		financial.NextYearForecastEarningsPerShare,
		financial.NumberOfIssuedAndOutstandingSharesAtTheEndOfFiscalYearIncludingTreasuryStock,
	)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

/*
財務情報テーブルにまとめて INSERT する関数
  - return) financial	財務情報
  - return) err			エラー
*/
func InsertFinancialInfoAll(financial []model.FinancialInfo) (err error) {
	// 財務情報テーブルに INSERT
	for _, state := range financial {
		_, err = db.Exec("INSERT INTO financial_info (code, disclosed_date, disclosed_time, net_sales, operating_profit, ordinary_profit, profit, earnings_per_share, total_assets, equity, equity_to_asset_ratio, book_value_per_share, cash_flows_from_operating_activities, cash_flows_from_investing_activities, cash_flows_from_financing_activities, cash_and_equivalents, result_dividend_per_share_annual, result_payout_ratio_annual, forecast_dividend_per_share_annual, forecast_payout_ratio_annual, next_year_forecast_dividend_per_share_annual, next_year_forecast_payout_ratio_annual, forecast_net_sales, forecast_operating_profit, forecast_ordinary_profit, forecast_profit, forecast_earnings_per_share, next_year_forecast_net_sales, next_year_forecast_operating_profit, next_year_forecast_ordinary_profit, next_year_forecast_profit, next_year_forecast_earnings_per_share, number_of_issued_and_outstanding_shares_at_the_end_of_fiscal_year_including_treasury_stock) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33)",
			state.Code,
			state.DisclosedDate,
			state.DisclosedTime,
			state.NetSales,
			state.OperatingProfit,
			state.OrdinaryProfit,
			state.Profit,
			state.EarningsPerShare,
			state.TotalAssets,
			state.Equity,
			state.EquityToAssetRatio,
			state.BookValuePerShare,
			state.CashFlowsFromOperatingActivities,
			state.CashFlowsFromInvestingActivities,
			state.CashFlowsFromFinancingActivities,
			state.CashAndEquivalents,
			state.ResultDividendPerShareAnnual,
			state.ResultPayoutRatioAnnual,
			state.ForecastDividendPerShareAnnual,
			state.ForecastPayoutRatioAnnual,
			state.NextYearForecastDividendPerShareAnnual,
			state.NextYearForecastPayoutRatioAnnual,
			state.ForecastNetSales,
			state.ForecastOperatingProfit,
			state.ForecastOrdinaryProfit,
			state.ForecastProfit,
			state.ForecastEarningsPerShare,
			state.NextYearForecastNetSales,
			state.NextYearForecastOperatingProfit,
			state.NextYearForecastOrdinaryProfit,
			state.NextYearForecastProfit,
			state.NextYearForecastEarningsPerShare,
			state.NumberOfIssuedAndOutstandingSharesAtTheEndOfFiscalYearIncludingTreasuryStock,
		)
	}
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

/*
財務情報テーブルを UPDATE する関数
  - arg) financial	財務情報
  - return) err		エラー
*/
func UpdateFinancialInfo(financial model.FinancialInfo) (err error) {
	// 財務情報テーブルを UPDATE
	_, err = db.Exec("UPDATE financial_info SET disclosed_date = $2, disclosed_time = $3, net_sales = $4, operating_profit = $5, ordinary_profit = $6, profit = $7, earnings_per_share = $8, total_assets = $9, equity = $10, equity_to_asset_ratio = $11, book_value_per_share = $12, cash_flows_from_operating_activities = $13, cash_flows_from_investing_activities = $14, cash_flows_from_financing_activities = $15, cash_and_equivalents = $16, result_dividend_per_share_annual = $17, result_payout_ratio_annual = $18, forecast_dividend_per_share_annual = $19, forecast_payout_ratio_annual = $20, next_year_forecast_dividend_per_share_annual = $21, next_year_forecast_payout_ratio_annual = $22, forecast_net_sales = $23, forecast_operating_profit = $24, forecast_ordinary_profit = $25, forecast_profit = $26, forecast_earnings_per_share = $27, next_year_forecast_net_sales = $28, next_year_forecast_operating_profit = $29, next_year_forecast_ordinary_profit = $30, next_year_forecast_profit = $31, next_year_forecast_earnings_per_share = $32, number_of_issued_and_outstanding_shares_at_the_end_of_fiscal_year_including_treasury_stock = $33 WHERE code = $1",
		financial.Code,
		financial.DisclosedDate,
		financial.DisclosedTime,
		financial.NetSales,
		financial.OperatingProfit,
		financial.OrdinaryProfit,
		financial.Profit,
		financial.EarningsPerShare,
		financial.TotalAssets,
		financial.Equity,
		financial.EquityToAssetRatio,
		financial.BookValuePerShare,
		financial.CashFlowsFromOperatingActivities,
		financial.CashFlowsFromInvestingActivities,
		financial.CashFlowsFromFinancingActivities,
		financial.CashAndEquivalents,
		financial.ResultDividendPerShareAnnual,
		financial.ResultPayoutRatioAnnual,
		financial.ForecastDividendPerShareAnnual,
		financial.ForecastPayoutRatioAnnual,
		financial.NextYearForecastDividendPerShareAnnual,
		financial.NextYearForecastPayoutRatioAnnual,
		financial.ForecastNetSales,
		financial.ForecastOperatingProfit,
		financial.ForecastOrdinaryProfit,
		financial.ForecastProfit,
		financial.ForecastEarningsPerShare,
		financial.NextYearForecastNetSales,
		financial.NextYearForecastOperatingProfit,
		financial.NextYearForecastOrdinaryProfit,
		financial.NextYearForecastProfit,
		financial.NextYearForecastEarningsPerShare,
		financial.NumberOfIssuedAndOutstandingSharesAtTheEndOfFiscalYearIncludingTreasuryStock,
	)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

/*
財務情報テーブルを取得する関数
  - arg) code			銘柄コード
  - return) financial	財務情報
  - return) err			エラー
*/
func GetFinancialInfo(code string) (financial model.FinancialInfo, err error) {
	// データの取得
	err = db.QueryRow("SELECT * FROM financial_info WHERE code = $1", code).Scan(
		&financial.Code,
		&financial.DisclosedDate,
		&financial.DisclosedTime,
		&financial.NetSales,
		&financial.OperatingProfit,
		&financial.OrdinaryProfit,
		&financial.Profit,
		&financial.EarningsPerShare,
		&financial.TotalAssets,
		&financial.Equity,
		&financial.EquityToAssetRatio,
		&financial.BookValuePerShare,
		&financial.CashFlowsFromOperatingActivities,
		&financial.CashFlowsFromInvestingActivities,
		&financial.CashFlowsFromFinancingActivities,
		&financial.CashAndEquivalents,
		&financial.ResultDividendPerShareAnnual,
		&financial.ResultPayoutRatioAnnual,
		&financial.ForecastDividendPerShareAnnual,
		&financial.ForecastPayoutRatioAnnual,
		&financial.NextYearForecastDividendPerShareAnnual,
		&financial.NextYearForecastPayoutRatioAnnual,
		&financial.ForecastNetSales,
		&financial.ForecastOperatingProfit,
		&financial.ForecastOrdinaryProfit,
		&financial.ForecastProfit,
		&financial.ForecastEarningsPerShare,
		&financial.NextYearForecastNetSales,
		&financial.NextYearForecastOperatingProfit,
		&financial.NextYearForecastOrdinaryProfit,
		&financial.NextYearForecastProfit,
		&financial.NextYearForecastEarningsPerShare,
		&financial.NumberOfIssuedAndOutstandingSharesAtTheEndOfFiscalYearIncludingTreasuryStock,
	)
	if err != nil {
		log.Error(err)
		return model.FinancialInfo{}, err
	}

	return financial, nil
}

/*
財務情報テーブルを取得する関数
  - return) financial	財務情報
  - return) err			エラー
*/
func GetFinancialInfoAll() (financial []model.FinancialInfo, err error) {
	// データの取得
	rows, err := db.Query("SELECT * FROM financial_info")
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// 取得したデータを格納
	for rows.Next() {
		var state model.FinancialInfo
		err := rows.Scan(
			&state.Code,
			&state.DisclosedDate,
			&state.DisclosedTime,
			&state.NetSales,
			&state.OperatingProfit,
			&state.OrdinaryProfit,
			&state.Profit,
			&state.EarningsPerShare,
			&state.TotalAssets,
			&state.Equity,
			&state.EquityToAssetRatio,
			&state.BookValuePerShare,
			&state.CashFlowsFromOperatingActivities,
			&state.CashFlowsFromInvestingActivities,
			&state.CashFlowsFromFinancingActivities,
			&state.CashAndEquivalents,
			&state.ResultDividendPerShareAnnual,
			&state.ResultPayoutRatioAnnual,
			&state.ForecastDividendPerShareAnnual,
			&state.ForecastPayoutRatioAnnual,
			&state.NextYearForecastDividendPerShareAnnual,
			&state.NextYearForecastPayoutRatioAnnual,
			&state.ForecastNetSales,
			&state.ForecastOperatingProfit,
			&state.ForecastOrdinaryProfit,
			&state.ForecastProfit,
			&state.ForecastEarningsPerShare,
			&state.NextYearForecastNetSales,
			&state.NextYearForecastOperatingProfit,
			&state.NextYearForecastOrdinaryProfit,
			&state.NextYearForecastProfit,
			&state.NextYearForecastEarningsPerShare,
			&state.NumberOfIssuedAndOutstandingSharesAtTheEndOfFiscalYearIncludingTreasuryStock,
		)
		if err != nil {
			return nil, err
		}
		financial = append(financial, state)
	}

	// エラーチェック
	if err = rows.Err(); err != nil {
		log.Error(err)
		return nil, err
	}

	return financial, nil
}

/*
財務情報テーブルを全て削除する関数
  - return) err		エラー
*/
func DeleteFinancialInfoAll() (err error) {
	// テーブルの全削除
	_, err = db.Exec("DELETE FROM financial_info")
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

/*
直近操作のあった行の数を取得する関数
  - return) rows	操作のあった行の数
  - return) err		エラー
*/
func RowsAffected() (rows int64, err error) {
    result, err := db.Exec("SELECT ROW_COUNT()")
    if err != nil {
        return 0, err
    }
    return result.RowsAffected()
}