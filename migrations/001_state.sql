do $$
begin
	if not exists (select * from pg_catalog.pg_type typ inner join pg_catalog.pg_namespace nsp on nsp."oid" = typ.typnamespace
		where nsp.nspname  = current_schema() and typ.typname = 'state') then
		CREATE TYPE "state" AS ENUM('1', '2');
		end if;
end ;
$$
