-- Verify 01-deployments

BEGIN;

SELECT id, sha, deployed_at
  FROM public.deployments
  WHERE FALSE;

ROLLBACK;
