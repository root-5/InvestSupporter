-- データベースの作成
CREATE DATABASE IF NOT EXISTS financial_data;
USE financial_data;

-- 1. 会社情報テーブル (companies_info)
CREATE TABLE companies_info (
    code CHAR(5) NOT NULL PRIMARY KEY,
    company_name VARCHAR(50),
    company_name_english VARCHAR(100),
    sector17_code SMALLINT,
    sector33_code SMALLINT,
    scale_category VARCHAR(50),
    market_code SMALLINT,
    margin_code SMALLINT,
    INDEX idx_code (code),
);

-- 2. 17 業種情報テーブル (sector17_info)
CREATE TABLE sector17_info (
    sector17_code SMALLINT NOT NULL PRIMARY KEY,
    sector17_name VARCHAR(50),
    INDEX idx_sector17_code (sector17_code)
);

-- 3. 33 業種情報テーブル (sector33_info)
CREATE TABLE sector33_info (
    sector33_code SMALLINT NOT NULL PRIMARY KEY,
    sector33_name VARCHAR(50),
    INDEX idx_sector33_code (sector33_code)
);

-- 4. 市場区分情報テーブル (market_info)
CREATE TABLE market_info (
    market_code SMALLINT NOT NULL PRIMARY KEY,
    market_name VARCHAR(50),
    INDEX idx_market_code (market_code)
);

-- 5. 貸借信用区分情報テーブル (margin_info)
CREATE TABLE margin_info (
    margin_code SMALLINT NOT NULL PRIMARY KEY,
    margin_name VARCHAR(50),
    INDEX idx_margin_code (margin_code)
);

-- 6. 財務情報テーブル (financial_info)
CREATE TABLE financial_info (
    code CHAR(5) NOT NULL,
    disclosed_date DATE,
    disclosed_time TIME,
    net_sales DECIMAL(20,0),
    operating_profit DECIMAL(20,0),
    ordinary_profit DECIMAL(20,0),
    profit DECIMAL(20,0),
    earnings_per_share DECIMAL(10,2),
    total_assets DECIMAL(20,0),
    equity DECIMAL(20,0),
    equity_to_asset_ratio DECIMAL(10,3),
    book_value_per_share DECIMAL(10,2),
    cash_flows_from_operating_activities DECIMAL(20,0),
    cash_flows_from_investing_activities DECIMAL(20,0),
    cash_flows_from_financing_activities DECIMAL(20,0),
    cash_and_equivalents DECIMAL(20,0),
    result_dividend_per_share_annual DECIMAL(10,2),
    result_payout_ratio_annual DECIMAL(10,3),
    forecast_dividend_per_share_annual DECIMAL(10,2),
    forecast_payout_ratio_annual DECIMAL(10,3),
    next_year_forecast_dividend_per_share_annual DECIMAL(10,2),
    next_year_forecast_payout_ratio_annual DECIMAL(10,3),
    forecast_net_sales DECIMAL(20,0),
    forecast_operating_profit DECIMAL(20,0),
    forecast_ordinary_profit DECIMAL(20,0),
    forecast_profit DECIMAL(20,0),
    forecast_earnings_per_share DECIMAL(10,2),
    next_year_forecast_net_sales DECIMAL(20,0),
    next_year_forecast_operating_profit DECIMAL(20,0),
    next_year_forecast_ordinary_profit DECIMAL(20,0),
    next_year_forecast_profit DECIMAL(20,0),
    next_year_forecast_earnings_per_share DECIMAL(10,2),
    number_of_issued_and_outstanding_shares_at_the_end_of_fiscal_year_including_treasury_stock DECIMAL(20,0),
    PRIMARY KEY (code, disclosed_date),
    INDEX idx_code (code),
    FOREIGN KEY (code) REFERENCES companies_info(code)
);

-- 7. 株価情報テーブル (price_info)
CREATE TABLE price_info (
    code CHAR(5) NOT NULL,
    date DATE NOT NULL,
    adjustment_open DECIMAL(10,2),
    adjustment_high DECIMAL(10,2),
    adjustment_low DECIMAL(10,2),
    adjustment_close DECIMAL(10,2),
    adjustment_volume DECIMAL(20,0),
    PRIMARY KEY (code, date),
    INDEX idx_code_date (code, date),
    FOREIGN KEY (code) REFERENCES companies_info(code)
);
