# ESO 包工头

## 这是干嘛的：

一个装修工具，可以从各个维度筛选家具，可以标记你已经有了的，或者你想要的，方便你猛猛装修
![image](https://github.com/user-attachments/assets/a198c02d-0e03-4d95-89c2-d276d8c135d3)


## 你这数据哪里来的

**ESO家具查询工具-20240712版**里解析的

## 装起来方便吗

不方便，需要有一定的计算机能力

## 现在做到哪一步了？

现在用go写了后端服务，写了一个html去展示筛选结果

但是分类筛选还没做好（大类小类联动还不知道怎么写）

分页还没做

## 那最终我要怎么用

0.把数据库搞好（安装pg，建表，初始化数据->在sql文件）

```sql
CREATE TABLE public.jj (
	f_id varchar(1000) NULL,
	f_name_en varchar(1000) NULL,
	f_name_zh varchar(1000) NULL,
	f_sub_id varchar(1000) NULL,
	station_zh varchar(1000) NULL,
	station_en varchar(1000) NULL,
	main_category_en varchar(1000) NULL,
	sub_category_en varchar(1000) NULL,
	main_category_zh varchar(1000) NULL,
	sub_category_zh varchar(1000) NULL,
	get_category_zh varchar(1000) NULL,
	best_get_way_zh varchar(1000) NULL,
	get_category_en varchar(1000) NULL,
	best_get_way_en varchar(1000) NULL,
	price_gold varchar(1000) NULL,
	price_crown varchar(1000) NULL,
	price_ap varchar(1000) NULL,
	price_master_writ varchar(1000) NULL,
	price_tel_stone varchar(1000) NULL,
	price_gem varchar(1000) NULL,
	price_end varchar(1000) NULL,
	recipe_name_en varchar(1000) NULL,
	recipe_material_en varchar(1000) NULL,
	recipe_name_zh varchar(1000) NULL,
	recipe_material_zh varchar(1000) NULL,
	furniture_image_complete varchar(1000) NULL,
	furniture_info_complete varchar(1000) NULL,
	furniture_image_local_path varchar(1000) NULL,
	tag1 varchar(1000) NULL,
	tag2 varchar(1000) NULL,
	tag3 varchar(1000) NULL,
	pic bytea NULL,
	do_i_have bool NULL,
	do_i_want bool NULL
);

CREATE INDEX jj_f_name_zh_idx ON public.jj (f_name_zh);
CREATE INDEX jj_tag1_idx ON public.jj (tag1);
CREATE INDEX jj_tag2_idx ON public.jj (tag2);
CREATE INDEX jj_tag3_idx ON public.jj (tag3);
CREATE INDEX jj_do_i_have_idx ON public.jj (do_i_have);
CREATE INDEX jj_do_i_want_idx ON public.jj (do_i_want);
CREATE INDEX jj_best_get_way_zh_idx ON public.jj (best_get_way_zh);
CREATE INDEX jj_main_category_zh_idx ON public.jj (main_category_zh);
CREATE INDEX jj_sub_category_zh_idx ON public.jj (sub_category_zh);
```

1.把exe跑起来，打开html用

2.把go文件自己跑起来，打开html用
