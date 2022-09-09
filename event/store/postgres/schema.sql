CREATE TABLE IF NOT EXISTS events (
    source_key UUID NOT NULL,
    source_sequence BIGINT NOT NULL,
    emit_timestamp TIMESTAMP WITH TIME ZONE NOT NULL,
    commit_timestamp TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    payload_type_url TEXT,
    payload_value BYTEA,
    PRIMARY KEY (source_key, source_sequence)
);

CREATE INDEX idx_events_emit_ts ON events USING btree (
    source_key ASC,
    emit_timestamp ASC
);

CREATE INDEX idx_events_store_ts ON events USING btree (
    source_key ASC,
    commit_timestamp ASC
);

CREATE TABLE IF NOT EXISTS snapshots (
    source_key UUID NOT NULL,
    source_sequence BIGINT NOT NULL,
    emit_timestamp TIMESTAMP WITH TIME ZONE NOT NULL,
    commit_timestamp TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    payload_type_url TEXT,
    payload_value BYTEA,
    PRIMARY KEY (source_key, source_sequence)
);

CREATE INDEX idx_snapshots_emit_ts ON snapshots USING btree (
    source_key ASC,
    emit_timestamp ASC
);

CREATE INDEX idx_snapshots_commit_ts ON snapshots USING btree (
    source_key ASC,
    commit_timestamp ASC
);