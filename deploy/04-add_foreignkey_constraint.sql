-- Deploy 04-add_foreignkey_constraint
-- requires: 03-add_project_id_to_projects

BEGIN;

SET client_min_messages = 'warning';

ALTER TABLE public.deployments
  ADD CONSTRAINT projectfk FOREIGN KEY (project_id) REFERENCES projects (id) MATCH FULL;

COMMIT;
