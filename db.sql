

--  {
--  	"level": "error",
--  	"message": "Failed to connect to DB",
--      "resourceId": "server-1234",
--  	"timestamp": "2023-09-15T08:00:00Z",
--  	"traceId": "abc-xyz-123",
--      "spanId": "span-456",
--      "commit": "5e5342f",
--      "metadata": {
--          "parentResourceId": "server-0987"
--      }
-- }

CREATE TABLE logs  (
    id SERIAL PRIMARY KEY,
    level VARCHAR(10) NOT NULL,
    message TEXT NOT NULL,
    resource_id VARCHAR(255) NOT NULL,
    timestamp TIMESTAMP NOT NULL,
    trace_id VARCHAR(255) NOT NULL,
    span_id VARCHAR(255) NOT NULL,
    commit VARCHAR(255) NOT NULL,
    metadata JSON NOT NULL
);
