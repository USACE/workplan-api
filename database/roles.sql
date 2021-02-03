-- Users and Roles for HHD workplan Webapp

-- User workplan_user
-- Note: Substitute real password for 'password'
CREATE USER workplan_user WITH ENCRYPTED PASSWORD 'password';
CREATE ROLE workplan_reader;
CREATE ROLE workplan_writer;
CREATE ROLE postgis_reader;

--------------------------------------------------------------------------
-- NOTE: IF USERS ALREADY EXIST ON DATABASE, JUST RUN FROM THIS POINT DOWN
--------------------------------------------------------------------------

-- Role workplan_reader
-- Tables specific to workplan app
GRANT SELECT ON
    employee,
    project,
    project_funding_realitycheck,
    timeperiod,
    commitment,
    purchase,
    contract,
    travel,
    leave,
    v_commitment,
    v_leave,
    V_project
TO workplan_reader;

-- Role workplan_writer
-- Tables specific to workplan app
GRANT INSERT,UPDATE,DELETE ON
    employee,
    project,
    project_funding_realitycheck,
    timeperiod,
    commitment,
    purchase,
    contract,
    travel,
    leave
TO workplan_writer;

-- Role postgis_reader
GRANT SELECT ON geometry_columns TO postgis_reader;
GRANT SELECT ON geography_columns TO postgis_reader;
GRANT SELECT ON spatial_ref_sys TO postgis_reader;
-- Grant Permissions to instrument_user
GRANT postgis_reader TO workplan_user;
GRANT workplan_reader TO workplan_user;
GRANT workplan_writer TO workplan_user;
