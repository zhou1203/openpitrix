CREATE TABLE IF NOT EXISTS task (
	task_id     VARCHAR(50)  NOT NULL,
	job_id      VARCHAR(50)  NOT NULL,
	task_action VARCHAR(50)  NOT NULL,
	status      VARCHAR(50)  NOT NULL,
	directive   TEXT         NOT NULL,
	create_time TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
	status_time TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
	error_code  INT(11)      NOT NULL,
	owner       VARCHAR(50) NOT NULL,
	executor    VARCHAR(50)  NOT NULL,
	node_id     VARCHAR(50)  NOT NULL,
	target      VARCHAR(50)  NOT NULL,
	PRIMARY KEY (task_id),
	INDEX task_job_id_index (job_id ASC),
	INDEX task_create_time_index (create_time ASC),
	INDEX task_task_action_index (task_action ASC),
	INDEX task_owner_index (owner ASC),
	INDEX task_status_index (status ASC),
	INDEX task_node_id_index (node_id ASC),
	INDEX task_target_index (target ASC)
);
