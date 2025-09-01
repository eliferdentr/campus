CREATE TYPE note_status AS ENUM (
    'pending',
    'approved',
    'rejected',
    'active',
    'inactive',
    'completed',
    'failed',
    'cancelled'
);