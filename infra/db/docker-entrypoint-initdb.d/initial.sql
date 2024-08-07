-- 1. 上場銘柄テーブル (stocks_info)
CREATE TABLE stocks_info (
    code CHAR(5) PRIMARY KEY,
    company_name VARCHAR(50),
    company_name_english VARCHAR(100),
    sector17_code SMALLINT,
    sector33_code SMALLINT,
    scale_category VARCHAR(50),
    market_code SMALLINT,
    margin_code SMALLINT
    FOREIGN KEY (sector17_code) REFERENCES sector17_info(sector17_code),
    FOREIGN KEY (sector33_code) REFERENCES sector33_info(sector33_code),
    FOREIGN KEY (market_code) REFERENCES market_info(market_code),
    FOREIGN KEY (margin_code) REFERENCES margin_info(margin_code)
);

-- 2. 17 業種情報テーブル (sector17_info)
CREATE TABLE sector17_info (
    sector17_code SMALLINT PRIMARY KEY,
    sector17_name VARCHAR(50)
);

-- 3. 33 業種情報テーブル (sector33_info)
CREATE TABLE sector33_info (
    sector33_code SMALLINT PRIMARY KEY,
    sector33_name VARCHAR(50)
);

-- 4. 市場区分情報テーブル (market_info)
CREATE TABLE market_info (
    market_code SMALLINT PRIMARY KEY,
    market_name VARCHAR(50)
);

-- 5. 貸借信用区分情報テーブル (margin_info)
CREATE TABLE margin_info (
    margin_code SMALLINT PRIMARY KEY,
    margin_name VARCHAR(50)
);

-- 6. 財務情報テーブル (financial_info)
CREATE TABLE financial_info (
    code CHAR(5) PRIMARY KEY,
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
    FOREIGN KEY (code) REFERENCES stocks_info(code)
);

-- 7. 株価情報テーブル (price_info)
CREATE TABLE price_info (
    code CHAR(5) PRIMARY KEY,
    ymd DATE,
    adjustment_open DECIMAL(10,2),
    adjustment_high DECIMAL(10,2),
    adjustment_low DECIMAL(10,2),
    adjustment_close DECIMAL(10,2),
    adjustment_volume DECIMAL(20,0),
    FOREIGN KEY (code) REFERENCES stocks_info(code)
);

-- インデックスの作成
CREATE INDEX idx_code_stocks_info ON stocks_info (code);
CREATE INDEX idx_code_financial_info ON financial_info (code);
CREATE INDEX idx_code_price_info ON price_info (code);

-- csv ファイルの読み込み
COPY sector17_info FROM '/docker-entrypoint-initdb.d/sector17_info.csv' WITH (FORMAT csv, HEADER true);
COPY sector33_info FROM '/docker-entrypoint-initdb.d/sector33_info.csv' WITH (FORMAT csv, HEADER true);
COPY market_info FROM '/docker-entrypoint-initdb.d/market_info.csv' WITH (FORMAT csv, HEADER true);
COPY margin_info FROM '/docker-entrypoint-initdb.d/margin_info.csv' WITH (FORMAT csv, HEADER true);
