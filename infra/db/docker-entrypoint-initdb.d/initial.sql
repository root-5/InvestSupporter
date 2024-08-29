-- 1. 17 業種情報テーブル (sector17_info)
CREATE TABLE sector17_info (
    sector17_code SMALLINT PRIMARY KEY,
    sector17_name VARCHAR(50)
);

-- 2. 33 業種情報テーブル (sector33_info)
CREATE TABLE sector33_info (
    sector33_code SMALLINT PRIMARY KEY,
    sector33_name VARCHAR(50)
);

-- 3. 市場区分情報テーブル (market_info)
CREATE TABLE market_info (
    market_code SMALLINT PRIMARY KEY,
    market_name VARCHAR(50)
);

-- 4. 上場銘柄テーブル (stocks_info)
CREATE TABLE stocks_info (
    code CHAR(5) PRIMARY KEY,
    company_name VARCHAR(100),
    company_name_english VARCHAR(200),
    sector17_code SMALLINT,
    sector33_code SMALLINT,
    scale_category VARCHAR(100),
    market_code SMALLINT,
    FOREIGN KEY (sector17_code) REFERENCES sector17_info(sector17_code),
    FOREIGN KEY (sector33_code) REFERENCES sector33_info(sector33_code),
    FOREIGN KEY (market_code) REFERENCES market_info(market_code)
);

-- 5. 財務情報テーブル (financial_info)
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

-- 6. 株価情報テーブル (price_info)
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

-- データの挿入
-- 17 業種情報テーブル
INSERT INTO sector17_info (sector17_code, sector17_name) VALUES
    (1,'食品'),
    (2,'エネルギー資源'),
    (3,'建設・資材'),
    (4,'素材・化学'),
    (5,'医薬品'),
    (6,'自動車・輸送機'),
    (7,'鉄鋼・非鉄'),
    (8,'機械'),
    (9,'電機・精密'),
    (10,'情報通信・サービスその他'),
    (11,'電気・ガス'),
    (12,'運輸・物流'),
    (13,'商社・卸売'),
    (14,'小売'),
    (15,'銀行'),
    (16,'金融（除く銀行）'),
    (17,'不動産'),
    (99,'その他');

-- 33 業種情報テーブル
INSERT INTO sector33_info (sector33_code, sector33_name) VALUES
    (50,'水産・農林業'),
    (1050,'鉱業'),
    (2050,'建設業'),
    (3050,'食料品'),
    (3100,'繊維製品'),
    (3150,'パルプ・紙'),
    (3200,'化学'),
    (3250,'医薬品'),
    (3300,'石油･石炭製品'),
    (3350,'ゴム製品'),
    (3400,'ガラス･土石製品'),
    (3450,'鉄鋼'),
    (3500,'非鉄金属'),
    (3550,'金属製品'),
    (3600,'機械'),
    (3650,'電気機器'),
    (3700,'輸送用機器'),
    (3750,'精密機器'),
    (3800,'その他製品'),
    (4050,'電気･ガス業'),
    (5050,'陸運業'),
    (5100,'海運業'),
    (5150,'空運業'),
    (5200,'倉庫･運輸関連業'),
    (5250,'情報･通信業'),
    (6050,'卸売業'),
    (6100,'小売業'),
    (7050,'銀行業'),
    (7100,'証券･商品先物取引業'),
    (7150,'保険業'),
    (7200,'その他金融業'),
    (8050,'不動産業'),
    (9050,'サービス業'),
    (9999,'その他');

-- 市場区分情報テーブル
INSERT INTO market_info (market_code, market_name) VALUES
    (101,'東証一部'),
    (102,'東証二部'),
    (104,'マザーズ'),
    (105,'TOKYO PRO MARKET'),
    (106,'JASDAQ スタンダード'),
    (107,'JASDAQ グロース'),
    (109,'その他'),
    (111,'プライム'),
    (112,'スタンダード'),
    (113,'グロース');
