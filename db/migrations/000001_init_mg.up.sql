begin;

drop type if exists team_status;
create type team_status as enum (
    'winning',
    'loosing',
    'prepare'
);

create table if not exists "team"(
    id uuid primary key default gen_random_uuid(),
    name varchar(50) not null,
    division_id uuid not null,
    play_off_id uuid,
    rating integer not null default 0,
    team_status team_status not null default 'prepare'
);

create table if not exists "division"(
    id uuid primary key default gen_random_uuid(),
    name varchar(2) not null,
    tournament_id uuid not null
);

create table if not exists "play_off"(
    id uuid primary key default gen_random_uuid(),
    tournament_id uuid not null,
    winner uuid
);

create table if not exists "tournament"(
    id uuid primary key default gen_random_uuid(),
    winner uuid
);

create table if not exists "match"(
    id uuid primary key default gen_random_uuid(),
    passed boolean not null default false,
    first_team_id uuid not null,
    second_team_id uuid not null,
    goals_first_team integer not null default 0,
    goals_second_team integer not null default 0,
    iteration int not null default 0,
    constraint fk_first_team_id foreign key (first_team_id) references team(id),
    constraint fk_second_team_id foreign key (second_team_id) references team(id)
);

create table if not exists "match_division"(
    match_id uuid not null,
    division_id uuid not null,
    constraint fk_match_id foreign key (match_id) references match(id),
    constraint fk_division_id foreign key (division_id) references division(id),
    constraint pk_match_division primary key (match_id, division_id)
);

create table if not exists "match_play_off"(
    match_id uuid not null,
    play_off_id uuid not null,
    constraint fk_match_id foreign key (match_id) references match(id),
    constraint fk_play_off_id foreign key (play_off_id) references play_off(id),
    constraint pk_match_play_off primary key (match_id, play_off_id)
);

alter table "team" add constraint fk_division_id foreign key (division_id) references division(id);
alter table "team" add constraint fk_play_off_id foreign key (play_off_id) references play_off(id);
alter table "division" add constraint fk_tournament_id foreign key (tournament_id) references tournament(id);
alter table "play_off" add constraint fk_tournament_id foreign key (tournament_id) references tournament(id);

end;