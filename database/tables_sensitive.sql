-- Queries to initialize data where values are not stored in tables.sql
-- to avoid saving plain-text on Github

-- Initialize Departmental Overhead and GA Overhead
-- Note: Replace 10.00 with actual values...
UPDATE timeperiod SET ohrate_department = 10.00, ohrate_ga = 10.00;

-- Update for timeperiods after a certain date
UPDATE timeperiod
SET ohrate_department = 10.00, ohrate_ga = 10.00
WHERE timeperiod_end > '2020-10-25 00:00:00+00'
