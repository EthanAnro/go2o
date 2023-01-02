/** 2023-01-02 订单返利 */
CREATE TABLE order_rebate_list (
  id              BIGSERIAL NOT NULL, 
  plan_id        int4 NOT NULL, 
  trader_id      int8 NOT NULL, 
  affiliate_code varchar(20) NOT NULL, 
  order_no       varchar(20) NOT NULL, 
  order_subject  varchar(40) NOT NULL, 
  order_amount   int8 NOT NULL, 
  rebase_amount  int8 NOT NULL, 
  status         int4 NOT NULL, 
  create_time    int8 NOT NULL, 
  update_time    int8 NOT NULL, 
  PRIMARY KEY (id));
COMMENT ON TABLE order_rebate_list IS '订单返利';
COMMENT ON COLUMN order_rebate_list.id IS '编号';
COMMENT ON COLUMN order_rebate_list.plan_id IS '返利方案Id';
COMMENT ON COLUMN order_rebate_list.trader_id IS '成交人Id';
COMMENT ON COLUMN order_rebate_list.affiliate_code IS '分享码';
COMMENT ON COLUMN order_rebate_list.order_no IS '订单号';
COMMENT ON COLUMN order_rebate_list.order_subject IS '订单标题';
COMMENT ON COLUMN order_rebate_list.order_amount IS '订单金额';
COMMENT ON COLUMN order_rebate_list.rebase_amount IS '返利金额';
COMMENT ON COLUMN order_rebate_list.status IS '返利状态';
COMMENT ON COLUMN order_rebate_list.create_time IS '创建时间';
COMMENT ON COLUMN order_rebate_list.update_time IS '更新时间';


CREATE TABLE order_rebate_item (
  id             SERIAL NOT NULL, 
  debate_id     int8 NOT NULL, 
  item_id       int8 NOT NULL, 
  item_name     varchar(20) NOT NULL, 
  item_image    varchar(120) NOT NULL, 
  item_amount   int8 NOT NULL, 
  rebate_amount int8 NOT NULL, 
  PRIMARY KEY (id));
COMMENT ON TABLE order_rebate_item IS '订单返利详情';
COMMENT ON COLUMN order_rebate_item.id IS '编号';
COMMENT ON COLUMN order_rebate_item.debate_id IS '返利单Id';
COMMENT ON COLUMN order_rebate_item.item_id IS '商品编号';
COMMENT ON COLUMN order_rebate_item.item_name IS '商品名称';
COMMENT ON COLUMN order_rebate_item.item_image IS '商品图片';
COMMENT ON COLUMN order_rebate_item.item_amount IS '商品金额';
COMMENT ON COLUMN order_rebate_item.rebate_amount IS '返利金额';