// PostgreSQL を利用するための関数をまとめたパッケージ
package postgres

import (
	"app/controller/log"
	"app/domain/model"

	_ "github.com/lib/pq"
)

/*
上場銘柄テーブルと財務情報デーブルを結合して返却する関数
  - return: 上場銘柄テーブルと財務情報デーブルを結合したデータ
*/
func GetFinancialInfoForApi() (financialInfoForApi []model.FinancialInfoForApi, err error) {
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
			financial.code,
			financial.disclosed_date,
			financial.disclosed_time,
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
		return financialInfoForApi, err
	}

	// データベースから取得したデータをスライスに格納
	for rows.Next() {
		var listedStockAndFinancialStatement model.FinancialInfoForApi
		err = rows.Scan(
			&listedStockAndFinancialStatement.StocksInfo.Code,
			&listedStockAndFinancialStatement.StocksInfo.CompanyName,
			&listedStockAndFinancialStatement.StocksInfo.CompanyNameEnglish,
			&listedStockAndFinancialStatement.StocksInfo.Sector17Code,
			&listedStockAndFinancialStatement.StocksInfo.Sector33Code,
			&listedStockAndFinancialStatement.StocksInfo.ScaleCategory,
			&listedStockAndFinancialStatement.StocksInfo.MarketCode,
			&listedStockAndFinancialStatement.FinancialInfo.Code,
			&listedStockAndFinancialStatement.FinancialInfo.DisclosedDate,
			&listedStockAndFinancialStatement.FinancialInfo.DisclosedTime,
			&listedStockAndFinancialStatement.FinancialInfo.NetSales,
			&listedStockAndFinancialStatement.FinancialInfo.OperatingProfit,
			&listedStockAndFinancialStatement.FinancialInfo.OrdinaryProfit,
			&listedStockAndFinancialStatement.FinancialInfo.Profit,
			&listedStockAndFinancialStatement.FinancialInfo.EarningsPerShare,
			&listedStockAndFinancialStatement.FinancialInfo.TotalAssets,
			&listedStockAndFinancialStatement.FinancialInfo.Equity,
			&listedStockAndFinancialStatement.FinancialInfo.EquityToAssetRatio,
			&listedStockAndFinancialStatement.FinancialInfo.BookValuePerShare,
			&listedStockAndFinancialStatement.FinancialInfo.CashFlowsFromOperatingActivities,
			&listedStockAndFinancialStatement.FinancialInfo.CashFlowsFromInvestingActivities,
			&listedStockAndFinancialStatement.FinancialInfo.CashFlowsFromFinancingActivities,
			&listedStockAndFinancialStatement.FinancialInfo.CashAndEquivalents,
			&listedStockAndFinancialStatement.FinancialInfo.ResultDividendPerShareAnnual,
			&listedStockAndFinancialStatement.FinancialInfo.ResultPayoutRatioAnnual,
			&listedStockAndFinancialStatement.FinancialInfo.ForecastDividendPerShareAnnual,
			&listedStockAndFinancialStatement.FinancialInfo.ForecastPayoutRatioAnnual,
			&listedStockAndFinancialStatement.FinancialInfo.NextYearForecastDividendPerShareAnnual,
			&listedStockAndFinancialStatement.FinancialInfo.NextYearForecastPayoutRatioAnnual,
			&listedStockAndFinancialStatement.FinancialInfo.ForecastNetSales,
			&listedStockAndFinancialStatement.FinancialInfo.ForecastOperatingProfit,
			&listedStockAndFinancialStatement.FinancialInfo.ForecastOrdinaryProfit,
			&listedStockAndFinancialStatement.FinancialInfo.ForecastProfit,
			&listedStockAndFinancialStatement.FinancialInfo.ForecastEarningsPerShare,
			&listedStockAndFinancialStatement.FinancialInfo.NextYearForecastNetSales,
			&listedStockAndFinancialStatement.FinancialInfo.NextYearForecastOperatingProfit,
			&listedStockAndFinancialStatement.FinancialInfo.NextYearForecastOrdinaryProfit,
			&listedStockAndFinancialStatement.FinancialInfo.NextYearForecastProfit,
			&listedStockAndFinancialStatement.FinancialInfo.NextYearForecastEarningsPerShare,
			&listedStockAndFinancialStatement.FinancialInfo.NumberOfIssuedAndOutstandingSharesAtTheEndOfFiscalYearIncludingTreasuryStock,
		)
		if err != nil {
			log.Error(err)
			return financialInfoForApi, err
		}
		financialInfoForApi = append(financialInfoForApi, listedStockAndFinancialStatement)
	}

	return financialInfoForApi, nil
}
