// PostgreSQL を利用するための関数をまとめたパッケージ
package postgres

import (
	"app/controller/log"
	"app/domain/model"
)

/*
上場銘柄テーブルと財務情報デーブルを結合して返却する関数
  - arg: 銘柄コード
  - return: 上場銘柄テーブルと財務情報デーブルを結合したデータ
*/
func GetFinancialInfoForApi(code string) (financialsInfoForApi []model.FinancialInfoForApi, err error) {
	// データベースからデータを取得
	rows, err := db.Query(`
		SELECT
			stocks.code,
			stocks.company_name,
			stocks.company_name_english,
			stocks.sector17_code,
			stocks.sector33_code,
			stocks.scale_category,
			stocks.market_code,
			financial.disclosed_date,
			financial.net_sales,
			financial.operating_profit,
			financial.ordinary_profit,
			financial.profit,
			financial.earnings_per_share,
			financial.total_assets,
			financial.equity,
			financial.equity_to_asset_ratio,
			financial.book_value_per_share,
			financial.cash_flows_from_operating_activities,
			financial.cash_flows_from_investing_activities,
			financial.cash_flows_from_financing_activities,
			financial.cash_and_equivalents,
			financial.result_dividend_per_share_annual,
			financial.result_payout_ratio_annual,
			financial.forecast_dividend_per_share_annual,
			financial.forecast_payout_ratio_annual,
			financial.next_year_forecast_dividend_per_share_annual,
			financial.next_year_forecast_payout_ratio_annual,
			financial.forecast_net_sales,
			financial.forecast_operating_profit,
			financial.forecast_ordinary_profit,
			financial.forecast_profit,
			financial.forecast_earnings_per_share,
			financial.next_year_forecast_net_sales,
			financial.next_year_forecast_operating_profit,
			financial.next_year_forecast_ordinary_profit,
			financial.next_year_forecast_profit,
			financial.next_year_forecast_earnings_per_share,
			financial.number_of_issued_and_outstanding_shares_at_the_end_of_fiscal_year_including_treasury_stock
		FROM
			stocks_info stocks
		INNER JOIN
			financial_info financial
		ON
			stocks.code = financial.code
		WHERE
			stocks.code = $1
	`, code)
	if err != nil {
		log.Error(err)
		return financialsInfoForApi, err
	}

	for rows.Next() {
		// データベースから取得したデータをスライスに格納
		var financialInfoForApi model.FinancialInfoForApi
		err = rows.Scan(
			&financialInfoForApi.Code,
			&financialInfoForApi.CompanyName,
			&financialInfoForApi.CompanyNameEnglish,
			&financialInfoForApi.Sector17Code,
			&financialInfoForApi.Sector33Code,
			&financialInfoForApi.ScaleCategory,
			&financialInfoForApi.MarketCode,
			&financialInfoForApi.DisclosedDate,
			&financialInfoForApi.NetSales,
			&financialInfoForApi.OperatingProfit,
			&financialInfoForApi.OrdinaryProfit,
			&financialInfoForApi.Profit,
			&financialInfoForApi.EarningsPerShare,
			&financialInfoForApi.TotalAssets,
			&financialInfoForApi.Equity,
			&financialInfoForApi.EquityToAssetRatio,
			&financialInfoForApi.BookValuePerShare,
			&financialInfoForApi.CashFlowsFromOperatingActivities,
			&financialInfoForApi.CashFlowsFromInvestingActivities,
			&financialInfoForApi.CashFlowsFromFinancingActivities,
			&financialInfoForApi.CashAndEquivalents,
			&financialInfoForApi.ResultDividendPerShareAnnual,
			&financialInfoForApi.ResultPayoutRatioAnnual,
			&financialInfoForApi.ForecastDividendPerShareAnnual,
			&financialInfoForApi.ForecastPayoutRatioAnnual,
			&financialInfoForApi.NextYearForecastDividendPerShareAnnual,
			&financialInfoForApi.NextYearForecastPayoutRatioAnnual,
			&financialInfoForApi.ForecastNetSales,
			&financialInfoForApi.ForecastOperatingProfit,
			&financialInfoForApi.ForecastOrdinaryProfit,
			&financialInfoForApi.ForecastProfit,
			&financialInfoForApi.ForecastEarningsPerShare,
			&financialInfoForApi.NextYearForecastNetSales,
			&financialInfoForApi.NextYearForecastOperatingProfit,
			&financialInfoForApi.NextYearForecastOrdinaryProfit,
			&financialInfoForApi.NextYearForecastProfit,
			&financialInfoForApi.NextYearForecastEarningsPerShare,
			&financialInfoForApi.NumberOfIssuedAndOutstandingSharesAtTheEndOfFiscalYearIncludingTreasuryStock,
		)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		financialsInfoForApi = append(financialsInfoForApi, financialInfoForApi)
	}

	return financialsInfoForApi, nil
}

/*
上場銘柄テーブルと財務情報デーブルを結合して返却する関数
  - return: 上場銘柄テーブルと財務情報デーブルを結合したデータ
*/
func GetFinancialsInfoForApi() (financialsInfoForApi []model.FinancialInfoForApi, err error) {
	// データベースからデータを取得
	rows, err := db.Query(`
		SELECT
			stocks.code,
			stocks.company_name,
			stocks.company_name_english,
			stocks.sector17_code,
			stocks.sector33_code,
			stocks.scale_category,
			stocks.market_code,
			financial.disclosed_date,
			financial.net_sales,
			financial.operating_profit,
			financial.ordinary_profit,
			financial.profit,
			financial.earnings_per_share,
			financial.total_assets,
			financial.equity,
			financial.equity_to_asset_ratio,
			financial.book_value_per_share,
			financial.cash_flows_from_operating_activities,
			financial.cash_flows_from_investing_activities,
			financial.cash_flows_from_financing_activities,
			financial.cash_and_equivalents,
			financial.result_dividend_per_share_annual,
			financial.result_payout_ratio_annual,
			financial.forecast_dividend_per_share_annual,
			financial.forecast_payout_ratio_annual,
			financial.next_year_forecast_dividend_per_share_annual,
			financial.next_year_forecast_payout_ratio_annual,
			financial.forecast_net_sales,
			financial.forecast_operating_profit,
			financial.forecast_ordinary_profit,
			financial.forecast_profit,
			financial.forecast_earnings_per_share,
			financial.next_year_forecast_net_sales,
			financial.next_year_forecast_operating_profit,
			financial.next_year_forecast_ordinary_profit,
			financial.next_year_forecast_profit,
			financial.next_year_forecast_earnings_per_share,
			financial.number_of_issued_and_outstanding_shares_at_the_end_of_fiscal_year_including_treasury_stock
		FROM
			stocks_info stocks
		INNER JOIN
			financial_info financial
		ON
			stocks.code = financial.code
	`)
	if err != nil {
		log.Error(err)
		return financialsInfoForApi, err
	}

	// データベースから取得したデータをスライスに格納
	for rows.Next() {
		var financialInfoForApi model.FinancialInfoForApi
		err = rows.Scan(
			&financialInfoForApi.Code,
			&financialInfoForApi.CompanyName,
			&financialInfoForApi.CompanyNameEnglish,
			&financialInfoForApi.Sector17Code,
			&financialInfoForApi.Sector33Code,
			&financialInfoForApi.ScaleCategory,
			&financialInfoForApi.MarketCode,
			&financialInfoForApi.DisclosedDate,
			&financialInfoForApi.NetSales,
			&financialInfoForApi.OperatingProfit,
			&financialInfoForApi.OrdinaryProfit,
			&financialInfoForApi.Profit,
			&financialInfoForApi.EarningsPerShare,
			&financialInfoForApi.TotalAssets,
			&financialInfoForApi.Equity,
			&financialInfoForApi.EquityToAssetRatio,
			&financialInfoForApi.BookValuePerShare,
			&financialInfoForApi.CashFlowsFromOperatingActivities,
			&financialInfoForApi.CashFlowsFromInvestingActivities,
			&financialInfoForApi.CashFlowsFromFinancingActivities,
			&financialInfoForApi.CashAndEquivalents,
			&financialInfoForApi.ResultDividendPerShareAnnual,
			&financialInfoForApi.ResultPayoutRatioAnnual,
			&financialInfoForApi.ForecastDividendPerShareAnnual,
			&financialInfoForApi.ForecastPayoutRatioAnnual,
			&financialInfoForApi.NextYearForecastDividendPerShareAnnual,
			&financialInfoForApi.NextYearForecastPayoutRatioAnnual,
			&financialInfoForApi.ForecastNetSales,
			&financialInfoForApi.ForecastOperatingProfit,
			&financialInfoForApi.ForecastOrdinaryProfit,
			&financialInfoForApi.ForecastProfit,
			&financialInfoForApi.ForecastEarningsPerShare,
			&financialInfoForApi.NextYearForecastNetSales,
			&financialInfoForApi.NextYearForecastOperatingProfit,
			&financialInfoForApi.NextYearForecastOrdinaryProfit,
			&financialInfoForApi.NextYearForecastProfit,
			&financialInfoForApi.NextYearForecastEarningsPerShare,
			&financialInfoForApi.NumberOfIssuedAndOutstandingSharesAtTheEndOfFiscalYearIncludingTreasuryStock,
		)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		financialsInfoForApi = append(financialsInfoForApi, financialInfoForApi)
	}

	return financialsInfoForApi, nil
}
