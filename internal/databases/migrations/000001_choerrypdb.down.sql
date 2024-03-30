BEGIN;

-- Drop foreign keys
ALTER TABLE "two_factor_authentication" DROP CONSTRAINT IF EXISTS "two_factor_authentication_user_id_fkey";
ALTER TABLE "password_reset_tokens" DROP CONSTRAINT IF EXISTS "password_reset_tokens_user_id_fkey";
ALTER TABLE "audit_logs" DROP CONSTRAINT IF EXISTS "audit_logs_user_id_fkey";
ALTER TABLE "role_permissions" DROP CONSTRAINT IF EXISTS "role_permissions_role_id_fkey";
ALTER TABLE "role_permissions" DROP CONSTRAINT IF EXISTS "role_permissions_permission_id_fkey";
ALTER TABLE "user_roles" DROP CONSTRAINT IF EXISTS "user_roles_user_id_fkey";
ALTER TABLE "user_roles" DROP CONSTRAINT IF EXISTS "user_roles_role_id_fkey";

-- Drop tables
DROP TABLE IF EXISTS "two_factor_authentication";
DROP TABLE IF EXISTS "password_reset_tokens";
DROP TABLE IF EXISTS "audit_logs";
DROP TABLE IF EXISTS "role_permissions";
DROP TABLE IF EXISTS "permissions";
DROP TABLE IF EXISTS "user_roles";
DROP TABLE IF EXISTS "roles";
DROP TABLE IF EXISTS "users";

-- Drop sequence
DROP SEQUENCE IF EXISTS users_id_seq;

COMMIT;
