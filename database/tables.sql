-- extensions
CREATE extension IF NOT EXISTS "uuid-ossp";

-- drop tables if they already exist
drop table if exists 
    public.employee,
    public.project,
    public.timeperiod,
    public.commitment
    CASCADE;

-- Employee
CREATE TABLE IF NOT EXISTS public.employee (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    name TEXT,
    rate NUMERIC(5,2),
    availability_multiplier NUMERIC DEFAULT 1.0
);

-- Timeperiod
CREATE TABLE IF NOT EXISTS public.timeperiod (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    name TEXT,
    timeperiod_end TIMESTAMPTZ,
    ohrate_department NUMERIC(5,2),
    ohrate_ga NUMERIC(5,2),
    workdays INTEGER
);

-- Project
CREATE TABLE IF NOT EXISTS public.project (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    name TEXT,
    timeperiod_id_end UUID REFERENCES timeperiod(id),
    funding NUMERIC(12,2),
    feedback_enabled BOOLEAN DEFAULT FALSE
);

-- Project Funding Reality Check
CREATE TABLE IF NOT EXISTS public.project_funding_realitycheck (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    project_id UUID NOT NULL REFERENCES project(id),
    timeperiod_id UUID NOT NULL REFERENCES timeperiod(id),
    total NUMERIC(12,2) NOT NULL,
    CONSTRAINT project_timeperiod_unique UNIQUE (project_id, timeperiod_id)
);

CREATE TABLE IF NOT EXISTS public.purchase (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    name TEXT,
    cost NUMERIC(12,2),
    project_id UUID REFERENCES project(id),
    timeperiod_id UUID REFERENCES timeperiod(id)
);

CREATE TABLE IF NOT EXISTS public.contract (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    name TEXT,
    cost NUMERIC(12,2),
    project_id UUID REFERENCES project(id),
    timeperiod_id UUID REFERENCES timeperiod(id)
);

CREATE TABLE IF NOT EXISTS public.travel (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    name TEXT,
    cost NUMERIC(12,2),
    project_id UUID REFERENCES project(id),
    timeperiod_id UUID REFERENCES timeperiod(id)
);

-- Commitment
CREATE TABLE IF NOT EXISTS public.commitment (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    employee_id UUID REFERENCES employee(id),
    timeperiod_id UUID REFERENCES timeperiod(id),
    project_id UUID REFERENCES project(id),
    days INTEGER NOT NULL,
    CONSTRAINT employee_timeperiod_project_unique UNIQUE (employee_id, timeperiod_id, project_id)
);

-- Leave
CREATE TABLE IF NOT EXISTS public.leave (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    employee_id UUID REFERENCES employee(id),
    timeperiod_id UUID REFERENCES timeperiod(id),
    days INTEGER NOT NULL,
    CONSTRAINT employee_timeperiod_unique UNIQUE (employee_id, timeperiod_id)
);

-- 
-- VIEWS
-- 

CREATE OR REPLACE VIEW public.v_project AS (
    SELECT p.id                 AS id,
           p.name               AS name,
           p.timeperiod_id_end  AS timeperiod_id_end,
           p.feedback_enabled   AS feedback_enabled,
           p.funding            AS funding,
           RC.total             AS funds_remaining,
           RC.timeperiod_end    AS latest_reality_check
    FROM project p
    LEFT JOIN (
        SELECT DISTINCT ON (b.project_id),
                            t.id             AS timeperiod_id,
                            t.timeperiod_end AS reality_check_date,
                            rc.total         AS reality_check_total,
        FROM project_funding_realitycheck rc
        INNER JOIN timeperiod t ON t.id = rc.timeperiod_id
        GROUP BY rc.project_id
        ORDER BY rc.project_id, t.timeperiod_end DESC
    ) RC ON RC.project_id = p.id
);

CREATE OR REPLACE VIEW public.v_commitment AS (
        SELECT C.id as id,
            T.id as timeperiod_id,
            T.name as timeperiod_name,
            P.id as project_id,
            P.name as project_name,
            E.id as employee_id,
            E.name as employee_name,
            C.days as days,
            ((C.days * 8 * E.rate)*(100 + t.ohrate_department + t.ohrate_ga)/100)::numeric(12,2) as cost
        FROM commitment C
            INNER JOIN timeperiod T on T.id = C.timeperiod_id
            INNER JOIN project P on P.id = C.project_id
            INNER JOIN employee E on E.id = C.employee_id
        ORDER BY employee_id,
            timeperiod_id,
            project_id
    );

CREATE OR REPLACE VIEW public.v_leave AS (
    SELECT L.id as id,
        T.id as timeperiod_id,
        T.name as timeperiod_name,
        E.id as employee_id,
        E.name as employee_name,
        L.days as days
    FROM leave L
        INNER JOIN timeperiod T on T.id = L.timeperiod_id
        INNER JOIN employee E on E.id = L.employee_id
    ORDER BY employee_id, timeperiod_id
);

-- v_employee_commitment_summary
-- CREATE OR REPLACE VIEW public.v_employee_commitment_summary AS (
--         SELECT E.id as employee_id,
--             E.name as employee_name,
--             T.id as timeperiod_id,
--             T.name as timeperiod_name,
--             COALESCE(S.committed_days, 0) * E.rate * 8 AS timeperiod_cost,
--             COALESCE(S.committed_days, 0) AS days_committed,
--             CAST(E.availability_multiplier * T.workdays as INT) - COALESCE(S.committed_days, 0) AS days_free,
--             CAST(
--                 100.0 * COALESCE(s.committed_days, 0) / (E.availability_multiplier * T.workdays) AS INT
--             ) AS committed_percent
--         FROM employee AS E
--             CROSS JOIN timeperiod T
--             LEFT JOIN (
--                 SELECT employee_id,
--                     timeperiod_id,
--                     sum(days) as committed_days
--                 FROM commitment
--                 GROUP BY employee_id,
--                     timeperiod_id
--             ) S ON S.employee_id = E.id
--             AND S.timeperiod_id = T.id
--     );

-- -- v_project_commitment_summary
-- CREATE OR REPLACE VIEW public.v_project_commitment_summary AS (
--         SELECT P.id as project_id,
--             P.name as project_name,
--             T.id as timeperiod_id,
--             T.name as timeperiod_name,
--             T.timeperiod_end as timeperiod_end,
--     COALESCE(S.charges, 0) AS charges
-- FROM project AS P
--     CROSS JOIN timeperiod T
--     LEFT JOIN (
--         SELECT project_id,
--             timeperiod_id,
--             sum(days * 8 * E.rate) AS charges
--         FROM commitment AS C
--             INNER JOIN employee E ON E.id = C.employee_id
--         GROUP BY project_id,
--             timeperiod_id
--     ) S ON S.project_id = P.id
--     AND S.timeperiod_id = T.id
-- );
-- -- v_timeperiod_capacity
-- CREATE OR REPLACE VIEW public.v_timeperiod_capacity AS (
--         SELECT tp.id as timeperiod_id,
--             tp.name as timeperiod_name,
--             sum(employee.rate * 8 * tp.workdays) as total_capacity
--         FROM timeperiod as tp
--             CROSS JOIN employee
--         GROUP BY tp.id
--     );

-- -------
-- Domains
-- -------
-- instrument_type
INSERT INTO employee (name, rate) VALUES
    ('Employe 1', 50.87),
    ('Employee 2', 76.71),
    ('Employee 3', 82.97);

-- timeperiod
INSERT INTO timeperiod (name, timeperiod_end, workdays, ohrate_department, ohrate_ga) VALUES
    ('Oct10 2020', 'October 10, 2020', 10, 0, 0),
    ('Oct24 2020', 'October 24, 2020', 10, 0, 0),
    ('Nov07 2020', 'November 7, 2020', 10, 0, 0),
    ('Nov21 2020', 'November 21, 2020', 9, 0, 0),
    ('Dec05 2020', 'December 5, 2020', 9, 0, 0),
    ('Dec19 2020', 'December 19, 2020', 10, 0, 0),
    ('Jan02 2021', 'January 2, 2021', 8, 0, 0),
    ('Jan16 2021', 'January 16, 2021', 10, 0, 0),
    ('Jan30 2021', 'January 30, 2021', 9, 0, 0),
    ('Feb13 2021', 'February 13, 2021', 10, 0, 0),
    ('Feb27 2021', 'February 27, 2021', 9, 0, 0),
    ('Mar13 2021', 'March 13, 2021', 10, 0, 0),
    ('Mar27 2021', 'March 27, 2021', 10, 0, 0),
    ('Apr10 2021', 'April 10, 2021', 10, 0, 0),
    ('Apr24 2021', 'April 24, 2021', 10, 0, 0),
    ('May08 2021', 'May 8, 2021', 10, 0, 0),
    ('May22 2021', 'May 22, 2021', 10, 0, 0),
    ('Jun05 2021', 'June 5, 2021', 9, 0, 0),
    ('Jun19 2021', 'June 19, 2021', 10, 0, 0),
    ('Jul03 2021', 'July 3, 2021', 10, 0, 0),
    ('Jul17 2021', 'July 17, 2021', 9, 0, 0),
    ('Jul31 2021', 'July 31, 2021', 10, 0, 0),
    ('Aug14 2021', 'August 14, 2021', 10, 0, 0),
    ('Aug28 2021', 'August 28, 2021', 10, 0, 0),
    ('Sep11 2021', 'September 11, 2021', 9, 0, 0),
    ('Sep25 2021', 'September 25, 2021', 10, 0, 0);

-- projects
INSERT INTO project (name, funding) VALUES
    ('Cumulus', 180000.00),
    ('Access to Water', 300000.00),
    ('Midas', 30000.00);
