create table topic_detail
(
	topic_id int auto_increment,
	topic_desc text null,
	topic_input text null,
	topic_output text null,
	topic_demo_input text null,
	topic_demo_output text null,
	constraint topic_detail_id_uindex
		unique (topic_id)
);

alter table topic_detail
	add primary key (topic_id);

