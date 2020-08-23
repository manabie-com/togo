CREATE OR REPLACE FUNCTION truncate_tables()
	RETURNS void
	LANGUAGE plpgsql
 AS $function$
 DECLARE
		 statements CURSOR FOR
				 SELECT tablename FROM pg_tables
				 WHERE tablename <> 'gorp_migrations' AND schemaname = 'public';
 BEGIN
		 FOR stmt IN statements LOOP
				 EXECUTE 'TRUNCATE TABLE ' || quote_ident(stmt.tablename) || ' CASCADE;';
		 END LOOP;
 END;
 $function$;