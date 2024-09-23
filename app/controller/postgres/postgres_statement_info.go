// PostgreSQL を利用するための関数をまとめたパッケージ
package postgres

import (
	"app/controller/log"
	"app/domain/model"
	"database/sql"
)

/*
財務情報テーブルに INSERT する関数
  - return) statements	財務情報
  - return) err			エラー
*/
func InsertStatementsInfo(statements []model.StatementInfo) (err error) {
	// Prepare を利用して SQL 文を実行
	stmt, err := db.Prepare("INSERT INTO statements_info (disclosure_number, code, disclosed_date, type_of_document, net_sales, operating_profit, ordinary_profit, profit, earnings_per_share, total_assets, equity, equity_to_asset_ratio, book_value_per_share, cash_flows_from_operating_activities, cash_flows_from_investing_activities, cash_flows_from_financing_activities, cash_and_equivalents, result_dividend_per_share_annual, result_payout_ratio_annual, forecast_dividend_per_share_annual, forecast_payout_ratio_annual, next_year_forecast_dividend_per_share_annual, next_year_forecast_payout_ratio_annual, forecast_net_sales, forecast_operating_profit, forecast_ordinary_profit, forecast_profit, forecast_earnings_per_share, next_year_forecast_net_sales, next_year_forecast_operating_profit, next_year_forecast_ordinary_profit, next_year_forecast_profit, next_year_forecast_earnings_per_share, number_of_issued_and_outstanding_shares_at_the_end_of_fiscal_year_including_treasury_stock) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34)")
	if err != nil {
		log.Error(err)
		return err
	}
	defer stmt.Close()

	// 財務情報テーブルに INSERT
	for _, statement := range statements {
		_, err = stmt.Exec(
			statement.DisclosureNumber,
			statement.Code,
			statement.DisclosedDate,
			statement.TypeOfDocument,
			statement.NetSales,
			statement.OperatingProfit,
			statement.OrdinaryProfit,
			statement.Profit,
			statement.EarningsPerShare,
			statement.TotalAssets,
			statement.Equity,
			statement.EquityToAssetRatio,
			statement.BookValuePerShare,
			statement.CashFlowsFromOperatingActivities,
			statement.CashFlowsFromInvestingActivities,
			statement.CashFlowsFromFinancingActivities,
			statement.CashAndEquivalents,
			statement.ResultDividendPerShareAnnual,
			statement.ResultPayoutRatioAnnual,
			statement.ForecastDividendPerShareAnnual,
			statement.ForecastPayoutRatioAnnual,
			statement.NextYearForecastDividendPerShareAnnual,
			statement.NextYearForecastPayoutRatioAnnual,
			statement.ForecastNetSales,
			statement.ForecastOperatingProfit,
			statement.ForecastOrdinaryProfit,
			statement.ForecastProfit,
			statement.ForecastEarningsPerShare,
			statement.NextYearForecastNetSales,
			statement.NextYearForecastOperatingProfit,
			statement.NextYearForecastOrdinaryProfit,
			statement.NextYearForecastProfit,
			statement.NextYearForecastEarningsPerShare,
			statement.NumberOfIssuedAndOutstandingSharesAtTheEndOfFiscalYearIncludingTreasuryStock,
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
  - arg) statement	財務情報
  - return) result	更新結果
  - return) err		エラー
*/
func UpdateStatementsInfo(statements []model.StatementInfo) (result sql.Result, err error) {
	// Prepare を利用して SQL 文を実行
	stmt, err := db.Prepare("UPDATE statements_info SET disclosure_number = $2 , disclosed_date = $3 , type_of_document = $4 , net_sales = $5 , operating_profit = $6 , ordinary_profit = $7 , profit = $8 , earnings_per_share = $9 , total_assets = $10 , equity = $11 , equity_to_asset_ratio = $12 , book_value_per_share = $13 , cash_flows_from_operating_activities = $14 , cash_flows_from_investing_activities = $15 , cash_flows_from_financing_activities = $16 , cash_and_equivalents = $17 , result_dividend_per_share_annual = $18 , result_payout_ratio_annual = $19 , forecast_dividend_per_share_annual = $20 , forecast_payout_ratio_annual = $21 , next_year_forecast_dividend_per_share_annual = $22 , next_year_forecast_payout_ratio_annual = $23 , forecast_net_sales = $24 , forecast_operating_profit = $25 , forecast_ordinary_profit = $26 , forecast_profit = $27 , forecast_earnings_per_share = $28 , next_year_forecast_net_sales = $29 , next_year_forecast_operating_profit = $30 , next_year_forecast_ordinary_profit = $31 , next_year_forecast_profit = $32 , next_year_forecast_earnings_per_share = $33 , number_of_issued_and_outstanding_shares_at_the_end_of_fiscal_year_including_treasury_stock = $34  WHERE code = $1")
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer stmt.Close()

	// 財務情報テーブルを UPDATE
	for _, statement := range statements {
		result, err = stmt.Exec(
			statement.DisclosureNumber,
			statement.Code,
			statement.DisclosedDate,
			statement.TypeOfDocument,
			statement.NetSales,
			statement.OperatingProfit,
			statement.OrdinaryProfit,
			statement.Profit,
			statement.EarningsPerShare,
			statement.TotalAssets,
			statement.Equity,
			statement.EquityToAssetRatio,
			statement.BookValuePerShare,
			statement.CashFlowsFromOperatingActivities,
			statement.CashFlowsFromInvestingActivities,
			statement.CashFlowsFromFinancingActivities,
			statement.CashAndEquivalents,
			statement.ResultDividendPerShareAnnual,
			statement.ResultPayoutRatioAnnual,
			statement.ForecastDividendPerShareAnnual,
			statement.ForecastPayoutRatioAnnual,
			statement.NextYearForecastDividendPerShareAnnual,
			statement.NextYearForecastPayoutRatioAnnual,
			statement.ForecastNetSales,
			statement.ForecastOperatingProfit,
			statement.ForecastOrdinaryProfit,
			statement.ForecastProfit,
			statement.ForecastEarningsPerShare,
			statement.NextYearForecastNetSales,
			statement.NextYearForecastOperatingProfit,
			statement.NextYearForecastOrdinaryProfit,
			statement.NextYearForecastProfit,
			statement.NextYearForecastEarningsPerShare,
			statement.NumberOfIssuedAndOutstandingSharesAtTheEndOfFiscalYearIncludingTreasuryStock,
		)
		if err != nil {
			log.Error(err)
			return nil, err
		}
	}

	return result, nil
}

/*
銘柄コードを指定して、財務情報テーブルを取得する関数
  - arg) code			銘柄コード
  - return) statement	財務情報
  - return) err			エラー
*/
func GetStatementsInfo(code string) (statements []model.StatementInfo, err error) {
	// データの取得
	var rows *sql.Rows

	// サブクエリの中身は、同年の同じタイプの文書の中で最新の開示番号一覧を取得するもの
	// 目的は、財務データの出力内容を決算短信に絞ること（LIKE句の目的）と決算短信の修正があった場合に最新のものだけを取得する（JOIN句の目的）こと
	rows, err = db.Query(`
		SELECT t1.*
		FROM statements_info AS t1
		INNER JOIN (
			SELECT MAX(disclosure_number) FROM statements_info WHERE code = $1 AND type_of_document LIKE '%%FinancialStatements%%' GROUP BY CONCAT(EXTRACT(YEAR FROM disclosed_date), type_of_document)
		) AS t2
			ON t1.disclosure_number = t2.max
		ORDER BY t1.disclosure_number ASC
		`, code)
	if err != nil {
		log.Error(err)
		return []model.StatementInfo{}, err
	}

	// 取得したデータを格納
	for rows.Next() {
		var statement model.StatementInfo
		err := rows.Scan(
			&statement.DisclosureNumber,
			&statement.Code,
			&statement.DisclosedDate,
			&statement.TypeOfDocument,
			&statement.NetSales,
			&statement.OperatingProfit,
			&statement.OrdinaryProfit,
			&statement.Profit,
			&statement.EarningsPerShare,
			&statement.TotalAssets,
			&statement.Equity,
			&statement.EquityToAssetRatio,
			&statement.BookValuePerShare,
			&statement.CashFlowsFromOperatingActivities,
			&statement.CashFlowsFromInvestingActivities,
			&statement.CashFlowsFromFinancingActivities,
			&statement.CashAndEquivalents,
			&statement.ResultDividendPerShareAnnual,
			&statement.ResultPayoutRatioAnnual,
			&statement.ForecastDividendPerShareAnnual,
			&statement.ForecastPayoutRatioAnnual,
			&statement.NextYearForecastDividendPerShareAnnual,
			&statement.NextYearForecastPayoutRatioAnnual,
			&statement.ForecastNetSales,
			&statement.ForecastOperatingProfit,
			&statement.ForecastOrdinaryProfit,
			&statement.ForecastProfit,
			&statement.ForecastEarningsPerShare,
			&statement.NextYearForecastNetSales,
			&statement.NextYearForecastOperatingProfit,
			&statement.NextYearForecastOrdinaryProfit,
			&statement.NextYearForecastProfit,
			&statement.NextYearForecastEarningsPerShare,
			&statement.NumberOfIssuedAndOutstandingSharesAtTheEndOfFiscalYearIncludingTreasuryStock,
		)
		if err != nil {
			return nil, err
		}
		statements = append(statements, statement)
	}

	return statements, nil
}

/*
最新の財務情報テーブルを取得する関数
  - return) statements	財務情報
  - return) err			エラー
*/
func GetStatementInfoAll() (statements []model.StatementInfo, err error) {
	// データの取得
	rows, err := db.Query("SELECT t1.* FROM statements_info t1 INNER JOIN (SELECT code,MAX(disclosure_number) FROM statements_info WHERE type_of_document LIKE '%%FinancialStatements%%' GROUP BY code ORDER BY code) t2 ON t1.code = t2.code AND t1.disclosure_number = t2.max")
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// 取得したデータを格納
	for rows.Next() {
		var statement model.StatementInfo
		err := rows.Scan(
			&statement.DisclosureNumber,
			&statement.Code,
			&statement.DisclosedDate,
			&statement.TypeOfDocument,
			&statement.NetSales,
			&statement.OperatingProfit,
			&statement.OrdinaryProfit,
			&statement.Profit,
			&statement.EarningsPerShare,
			&statement.TotalAssets,
			&statement.Equity,
			&statement.EquityToAssetRatio,
			&statement.BookValuePerShare,
			&statement.CashFlowsFromOperatingActivities,
			&statement.CashFlowsFromInvestingActivities,
			&statement.CashFlowsFromFinancingActivities,
			&statement.CashAndEquivalents,
			&statement.ResultDividendPerShareAnnual,
			&statement.ResultPayoutRatioAnnual,
			&statement.ForecastDividendPerShareAnnual,
			&statement.ForecastPayoutRatioAnnual,
			&statement.NextYearForecastDividendPerShareAnnual,
			&statement.NextYearForecastPayoutRatioAnnual,
			&statement.ForecastNetSales,
			&statement.ForecastOperatingProfit,
			&statement.ForecastOrdinaryProfit,
			&statement.ForecastProfit,
			&statement.ForecastEarningsPerShare,
			&statement.NextYearForecastNetSales,
			&statement.NextYearForecastOperatingProfit,
			&statement.NextYearForecastOrdinaryProfit,
			&statement.NextYearForecastProfit,
			&statement.NextYearForecastEarningsPerShare,
			&statement.NumberOfIssuedAndOutstandingSharesAtTheEndOfFiscalYearIncludingTreasuryStock,
		)
		if err != nil {
			return nil, err
		}
		statements = append(statements, statement)
	}

	// エラーチェック
	if err = rows.Err(); err != nil {
		log.Error(err)
		return nil, err
	}

	return statements, nil
}

/*
財務情報テーブルを全て削除する関数
  - return) err		エラー
*/
func DeleteStatementInfoAll() (err error) {
	// テーブルの全削除
	_, err = db.Exec("DELETE FROM statements_info")
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
