-- name: GetUserByEmail :one
select
    id,
    email
from auth.users
where email = $1;

-- name: ListOAuthTokensByProvider :many
select
    user_id,
    provider,
    token,
    created_at,
    updated_at
from public.user_oauth_token
where provider = $1;

-- name: GetUserOAuthToken :one
select
    user_id,
    provider,
    token,
    created_at,
    updated_at
from public.user_oauth_token
where user_id = $1 and provider = $2;

-- name: GetUserEmailSyncHistory :one
select
    user_id,
    history_id,
    created_at,
    updated_at
from public.user_email_sync_history
where user_id = $1;

-- name: UspertUserEmailSyncHistory :exec
insert into public.user_email_sync_history(user_id, history_id)
values ($1, $2)
on conflict (user_id) do update set history_id = excluded.history_id;
