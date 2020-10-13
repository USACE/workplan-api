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
    rate REAL,
    availability_multiplier NUMERIC DEFAULT 1.0
);

-- Timeperiod
CREATE TABLE IF NOT EXISTS public.timeperiod (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    name TEXT,
    timeperiod_end TIMESTAMPTZ,
    workdays INTEGER
);

-- Project
CREATE TABLE IF NOT EXISTS public.project (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    name TEXT,
    timeperiod_id_end UUID REFERENCES timeperiod(id),
    funding INTEGER
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

-- 
-- VIEWS
-- 

CREATE OR REPLACE VIEW public.v_commitment AS (
        SELECT C.id as id,
            T.id as timeperiod_id,
            T.name as timeperiod_name,
            P.id as project_id,
            P.name as project_name,
            E.id as employee_id,
            E.name as employee_name,
            C.days as days,
            C.days * 8 * E.rate as cost
        FROM commitment C
            INNER JOIN timeperiod T on T.id = C.timeperiod_id
            INNER JOIN project P on P.id = C.project_id
            INNER JOIN employee E on E.id = C.employee_id
        ORDER BY employee_id,
            timeperiod_id,
            project_id
    );

-- v_employee_commitment_summary
CREATE OR REPLACE VIEW public.v_employee_commitment_summary AS (
        SELECT E.id as employee_id,
            E.name as employee_name,
            T.id as timeperiod_id,
            T.name as timeperiod_name,
            COALESCE(S.committed_days, 0) * E.rate * 8 AS timeperiod_cost,
            COALESCE(S.committed_days, 0) AS days_committed,
            CAST(E.availability_multiplier * T.workdays as INT) - COALESCE(S.committed_days, 0) AS days_free,
            CAST(
                100.0 * COALESCE(s.committed_days, 0) / (E.availability_multiplier * T.workdays) AS INT
            ) AS committed_percent
        FROM employee AS E
            CROSS JOIN timeperiod T
            LEFT JOIN (
                SELECT employee_id,
                    timeperiod_id,
                    sum(days) as committed_days
                FROM commitment
                GROUP BY employee_id,
                    timeperiod_id
            ) S ON S.employee_id = E.id
            AND S.timeperiod_id = T.id
    );

-- v_project_commitment_summary
CREATE OR REPLACE VIEW public.v_project_commitment_summary AS (
        SELECT P.id as project_id,
            P.name as project_name,
            T.id as timeperiod_id,
            T.name as timeperiod_name,
            T.timeperiod_end as timeperiod_end,
    COALESCE(S.charges, 0) AS charges
FROM project AS P
    CROSS JOIN timeperiod T
    LEFT JOIN (
        SELECT project_id,
            timeperiod_id,
            sum(days * 8 * E.rate) AS charges
        FROM commitment AS C
            INNER JOIN employee E ON E.id = C.employee_id
        GROUP BY project_id,
            timeperiod_id
    ) S ON S.project_id = P.id
    AND S.timeperiod_id = T.id
);
-- v_timeperiod_capacity
CREATE OR REPLACE VIEW public.v_timeperiod_capacity AS (
        SELECT tp.id as timeperiod_id,
            tp.name as timeperiod_name,
            sum(employee.rate * 8 * tp.workdays) as total_capacity
        FROM timeperiod as tp
            CROSS JOIN employee
        GROUP BY tp.id
    );
-- -------
-- Domains
-- -------
-- instrument_type
INSERT INTO employee (name, rate) VALUES
    ('Employe 1', 100),
    ('Employee 2', 120),
    ('Employee 3', 140);

-- timeperiod
INSERT INTO timeperiod (name, timeperiod_end, workdays) VALUES
    ('Oct10 2020', 'October 10, 2020', 10),
    ('Oct24 2020', 'October 24, 2020', 10),
    ('Nov07 2020', 'November 7, 2020', 10),
    ('Nov21 2020', 'November 21, 2020', 9),
    ('Dec05 2020', 'December 5, 2020', 9),
    ('Dec19 2020', 'December 19, 2020', 10),
    ('Jan02 2021', 'January 2, 2021', 8),
    ('Jan16 2021', 'January 16, 2021', 10),
    ('Jan30 2021', 'January 30, 2021', 9),
    ('Feb13 2021', 'February 13, 2021', 10),
    ('Feb27 2021', 'February 27, 2021', 9),
    ('Mar13 2021', 'March 13, 2021', 10),
    ('Mar27 2021', 'March 27, 2021', 10),
    ('Apr10 2021', 'April 10, 2021', 10),
    ('Apr24 2021', 'April 24, 2021', 10),
    ('May08 2021', 'May 8, 2021', 10),
    ('May22 2021', 'May 22, 2021', 10),
    ('Jun05 2021', 'June 5, 2021', 9),
    ('Jun19 2021', 'June 19, 2021', 10),
    ('Jul03 2021', 'July 3, 2021', 10),
    ('Jul17 2021', 'July 17, 2021', 9),
    ('Jul31 2021', 'July 31, 2021', 10),
    ('Aug14 2021', 'August 14, 2021', 10),
    ('Aug28 2021', 'August 28, 2021', 10),
    ('Sep11 2021', 'September 11, 2021', 9),
    ('Sep25 2021', 'September 25, 2021', 10);

-- projects
INSERT INTO project (name, funding) VALUES
    ('Cumulus', 180000),
    ('Access to Water', 300000),
    ('Midas', 30000);