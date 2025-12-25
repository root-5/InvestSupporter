package postgres

import (
	"app/controller/log"
	"app/domain/model"
)

/*
上場銘柄テーブルと最新の財務情報デーブルを結合してして返却する関数
  - return: 上場銘柄テーブルと財務情報デーブルを結合したデータ
*/
func GetBasicInfoForApi() (statementsInfoForApi []model.BasicInfoForApi, err error) {
	// データベースからデータを取得
	// LEFT JOIN で使用しているサブクエリは、財務情報テーブルの中で最新の財務情報を取得するためのクエリ
	rows, err := db.Query(`
		SELECT
			stocks.code,
			stocks.company_name,
			stocks.company_name_english,
			sector17.sector17_name,
			sector33.sector33_name,
			stocks.scale_category,
			market.market_name,
			statements.disclosed_date,
			statements.type_of_document,
			statements.net_sales,
			statements.operating_profit,
			statements.ordinary_profit,
			statements.profit,
			statements.earnings_per_share,
			statements.total_assets,
			statements.equity,
			statements.equity_to_asset_ratio,
			statements.book_value_per_share,
			statements.cash_flows_from_operating_activities,
			statements.cash_flows_from_investing_activities,
			statements.cash_flows_from_financing_activities,
			statements.cash_and_equivalents,
			statements.result_dividend_per_share_annual,
			statements.result_payout_ratio_annual,
			statements.forecast_dividend_per_share_annual,
			statements.next_year_forecast_dividend_per_share_annual,
			statements.next_year_forecast_payout_ratio_annual,
			statements.forecast_net_sales,
			statements.forecast_operating_profit,
			statements.forecast_ordinary_profit,
			statements.forecast_profit,
			statements.forecast_earnings_per_share,
			statements.next_year_forecast_net_sales,
			statements.next_year_forecast_operating_profit,
			statements.next_year_forecast_ordinary_profit,
			statements.next_year_forecast_profit,
			statements.next_year_forecast_earnings_per_share,
			statements.number_of_issued_and_outstanding_shares_at_the_end_of_fiscal_year_including_treasury_stock
		FROM
			stocks_info stocks
		LEFT JOIN (
			SELECT t1.* 
			FROM statements_info t1
			INNER JOIN (
				SELECT code,MAX(disclosure_number)
				FROM statements_info
				WHERE type_of_document
				LIKE '%%FinancialStatements%%'
				GROUP BY code ORDER BY code) t2
				ON t1.code = t2.code AND t1.disclosure_number = t2.max
			) statements
		ON
			stocks.code = statements.code
		LEFT JOIN
			sector17_info sector17
		ON
			stocks.sector17_code = sector17.sector17_code
		LEFT JOIN
			sector33_info sector33
		ON
			stocks.sector33_code = sector33.sector33_code
		LEFT JOIN
			markets_info market
		ON
			stocks.market_code = market.market_code
		WHERE
			sector33.sector33_name != 'その他'
		ORDER BY stocks.code
	`)
	if err != nil {
		log.Error(err)
		return statementsInfoForApi, err
	}

	// データベースから取得したデータをスライスに格納
	for rows.Next() {
		var statementInfoForApi model.BasicInfoForApi
		err = rows.Scan(
			&statementInfoForApi.Code,
			&statementInfoForApi.CompanyName,
			&statementInfoForApi.CompanyNameEnglish,
			&statementInfoForApi.Sector17Name,
			&statementInfoForApi.Sector33Name,
			&statementInfoForApi.ScaleCategory,
			&statementInfoForApi.MarketName,
			&statementInfoForApi.DisclosedDate,
			&statementInfoForApi.TypeOfDocument,
			&statementInfoForApi.NetSales,
			&statementInfoForApi.OperatingProfit,
			&statementInfoForApi.OrdinaryProfit,
			&statementInfoForApi.Profit,
			&statementInfoForApi.EarningsPerShare,
			&statementInfoForApi.TotalAssets,
			&statementInfoForApi.Equity,
			&statementInfoForApi.EquityToAssetRatio,
			&statementInfoForApi.BookValuePerShare,
			&statementInfoForApi.CashFlowsFromOperatingActivities,
			&statementInfoForApi.CashFlowsFromInvestingActivities,
			&statementInfoForApi.CashFlowsFromFinancingActivities,
			&statementInfoForApi.CashAndEquivalents,
			&statementInfoForApi.ResultDividendPerShareAnnual,
			&statementInfoForApi.ResultPayoutRatioAnnual,
			&statementInfoForApi.ForecastDividendPerShareAnnual,
			&statementInfoForApi.NextYearForecastDividendPerShareAnnual,
			&statementInfoForApi.NextYearForecastPayoutRatioAnnual,
			&statementInfoForApi.ForecastNetSales,
			&statementInfoForApi.ForecastOperatingProfit,
			&statementInfoForApi.ForecastOrdinaryProfit,
			&statementInfoForApi.ForecastProfit,
			&statementInfoForApi.ForecastEarningsPerShare,
			&statementInfoForApi.NextYearForecastNetSales,
			&statementInfoForApi.NextYearForecastOperatingProfit,
			&statementInfoForApi.NextYearForecastOrdinaryProfit,
			&statementInfoForApi.NextYearForecastProfit,
			&statementInfoForApi.NextYearForecastEarningsPerShare,
			&statementInfoForApi.NumberOfIssuedAndOutstandingSharesAtTheEndOfFiscalYearIncludingTreasuryStock,
		)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		statementsInfoForApi = append(statementsInfoForApi, statementInfoForApi)
	}

	return statementsInfoForApi, nil
}
