<div align="center">
    <img src=".github/banner.png" alt="Pocket Network logo" width="600"/>
    <!-- TODO Rename header -->
    <h1>Transaction DB</h1>
    <big>Driver implementation for the Transaction Database</big>
    <div>
    <br/>
        <a href="https://github.com/pokt-foundation/transaction-db/pulse"><img src="https://img.shields.io/github/last-commit/pokt-foundation/transaction-db.svg"/></a>
        <a href="https://github.com/pokt-foundation/transaction-db/pulls"><img src="https://img.shields.io/github/issues-pr/pokt-foundation/transaction-db.svg"/></a>
        <a href="https://github.com/pokt-foundation/transaction-db/issues"><img src="https://img.shields.io/github/issues-closed/pokt-foundation/transaction-db.svg"/></a>
    </div>
</div>
<br/>

  <!-- TODO Update the nelow section with development instructions (leave the pre-commit section in place) -->

# Development

## Pre-Commit Installation

Before starting development work on this repo, `pre-commit` must be installed.

In order to do so, run the command **`make init-pre-commit`** from the repository root.

Once this is done, the following checks will be performed on every commit to the repo and must pass before the commit is allowed:

### 1. Basic checks

- **check-yaml** - Checks YAML files for errors
- **check-merge-conflict** - Ensures there are no merge conflict markers
- **end-of-file-fixer** - Adds a newline to end of files
- **trailing-whitespace** - Trims trailing whitespace
- **no-commit-to-branch** - Ensures commits are not made directly to `main`

### 2. Go-specific checks

- **go-fmt** - Runs `gofmt`
- **go-imports** - Runs `goimports`
- **golangci-lint** - run `golangci-lint run ./...`
- **go-critic** - run `gocritic check ./...`
- **go-build** - run `go build`
- **go-mod-tidy** - run `go mod tidy -v`

### 3. Detect Secrets

Will detect any potential secrets or sensitive information before allowing a commit.

- Test variables that may resemble secrets (random hex strings, etc.) should be prefixed with `test_`
- The inline comment `pragma: allowlist secret` may be added to a line to force acceptance of a false positive

# Current Schema

```mermaid
erDiagram

  pocket_session {
    id bigint
    session_key char(44)
    session_height integer
    portal_region_name varchar
    created_at timestamp
    updated_at timestamp
  }

  portal_region {
    portal_region_name varchar
  }

  relay {
    id bigint
    pokt_chain_id char(4)
    endpoint_id varchar
    session_key char(44)
    protocol_app_public_key char(64)
    relay_source_url varchar
    pokt_node_address char(40)
    pokt_node_domain varchar
    pokt_node_public_key char(64)
    relay_start_datetime timestamp
    relay_return_datetime timestamp
    is_error boolean
    error_code integer
    error_name varchar
    error_message varchar
    error_source varchar
    error_type varchar
    relay_roundtrip_time float
    relay_chain_method_ids varchar
    relay_data_size integer
    relay_portal_trip_time float
    relay_node_trip_time float
    relay_url_is_public_endpoint boolean
    portal_region_name varchar
    is_altruist_relay boolean
    is_user_relay boolean
    request_id varchar
    pokt_tx_id varchar
    created_at timestamp
    updated_at timestamp
  }

  service_record {
    id bigint
    node_public_key char(64)
    pokt_chain_id char(4)
    session_key char(44)
    request_id varchar
    portal_region_name varchar
    latency float
    tickets integer
    result varchar
    available boolean
    successes integer
    failures integer
    p90_success_latency float
    median_success_latency float
    weighted_success_latency float
    success_rate float
    created_at timestamp
    updated_at timestamp
  }

  portal_region ||--o{ pocket_session : fk_pocket_session_portal_region
  portal_region ||--o{ relay : fk_relay_portal_region
  portal_region ||--o{ service_record : fk_service_region_portal_region

  pocket_session ||--o{ relay : fk_relay_session
  pocket_session ||--o{ service_record : fk_service_record_session
```