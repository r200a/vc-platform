CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
    user_id       UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email         VARCHAR(255) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    role          VARCHAR(20) NOT NULL CHECK (role IN ('founder', 'vc', 'admin')),
    refresh_token TEXT,
    is_verified   BOOLEAN DEFAULT FALSE,
    is_active     BOOLEAN DEFAULT TRUE,
    created_at    TIMESTAMP DEFAULT NOW(),
    updated_at    TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS user_profiles (
    profile_id    UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id       UUID UNIQUE NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    name          VARCHAR(255) NOT NULL,
    bio           TEXT,
    location      VARCHAR(255),
    profile_image TEXT,
    linkedin_url  VARCHAR(500),
    created_at    TIMESTAMP DEFAULT NOW(),
    updated_at    TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS vc_profiles (
    vc_id              UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id            UUID UNIQUE NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    fund_name          VARCHAR(255) NOT NULL,
    fund_size          BIGINT,
    ticket_size_min    BIGINT,
    ticket_size_max    BIGINT,
    focus_industries   TEXT[],
    focus_stages       TEXT[],
    website_url        VARCHAR(500),
    created_at         TIMESTAMP DEFAULT NOW(),
    updated_at         TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS startup_profiles (
    startup_id      UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    founder_id      UUID NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    name            VARCHAR(255) NOT NULL,
    tagline         VARCHAR(500),
    industry        VARCHAR(100),
    stage           VARCHAR(50) CHECK (stage IN ('idea','pre-seed','seed','series-a','series-b')),
    revenue_monthly BIGINT DEFAULT 0,
    pitch_deck_url  TEXT,
    website_url     VARCHAR(500),
    team_size       INTEGER DEFAULT 1,
    is_active       BOOLEAN DEFAULT TRUE,
    created_at      TIMESTAMP DEFAULT NOW(),
    updated_at      TIMESTAMP DEFAULT NOW()
);

DO $$ BEGIN
    CREATE TYPE application_status AS ENUM (
        'applied',
        'shortlisted',
        'pitching',
        'funded',
        'rejected'
    );
EXCEPTION
    WHEN duplicate_object THEN NULL;
END $$;

CREATE TABLE IF NOT EXISTS applications (
    application_id  UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    startup_id      UUID NOT NULL REFERENCES startup_profiles(startup_id) ON DELETE CASCADE,
    vc_id           UUID NOT NULL REFERENCES vc_profiles(vc_id) ON DELETE CASCADE,
    status          application_status DEFAULT 'applied',
    cover_note      TEXT,
    rejection_note  TEXT,
    applied_at      TIMESTAMP DEFAULT NOW(),
    updated_at      TIMESTAMP DEFAULT NOW(),
    UNIQUE(startup_id, vc_id)
);

-- indexes
CREATE INDEX IF NOT EXISTS idx_users_email          ON users(email);
CREATE INDEX IF NOT EXISTS idx_applications_vc      ON applications(vc_id);
CREATE INDEX IF NOT EXISTS idx_applications_startup ON applications(startup_id);
CREATE INDEX IF NOT EXISTS idx_applications_status  ON applications(status);