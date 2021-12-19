create table topic_list
(
	id int auto_increment,
	page_id int null,
	topic_id int null,
	topic_title text null,
	topic_href text null,
	topic_cat text null,
	topic_easy int null,
	constraint table_name_id_uindex
		unique (id)
);

alter table topic_list
	add primary key (id);

