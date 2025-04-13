#!/bin/sh

echo "Starting migrations..."

# https://docs.liquibase.com/concepts/changelogs/how-liquibase-finds-files.html

# In addition to JAR files, Liquibase also needs to find the changelog and looks for both of these file types in the following locations:

# Any paths specified in the classpath setting.
# The current working directory.
# A liquibase_libs directory in the current working directory.
# A lib directory in the Liquibase install location.
# Inside any .zip or .jar files in the liquibase_libs or lib directories.
# Using search-path (v4.13 and later). For more information about search-path, continue reading this page.

dbUrl="jdbc:postgresql://postgres:5432/backend_project"
dbUserName="developer"
dbPassword=""
defaultSchemaName="room_management"

export LIQUIBASE_HOME=/backend_project/liquibase

cd ${LIQUIBASE_HOME}

java -jar ${LIQUIBASE_HOME}/internal/lib/liquibase.jar \
    --driver=org.postgresql.Driver \
    --classpath=${LIQUIBASE_HOME}/internal/lib/postgresql.jar \
    --url=$dbUrl \
    --changeLogFile=./changelog/changelog-initschema.xml \
    --username=$dbUserName \
    --password=$dbPassword \
    --defaultSchemaName=public \
    --databaseChangeLogLockTableName="databasechangeloglock" \
    --databaseChangeLogTableName="databasechangelog" \
    update

java -jar ${LIQUIBASE_HOME}/internal/lib/liquibase.jar \
    --driver=org.postgresql.Driver \
    --classpath=${LIQUIBASE_HOME}/internal/lib/postgresql.jar \
    --url=$dbUrl \
    --changeLogFile=./changelog/changelog-master.xml \
    --username=$dbUserName \
    --password=$dbPassword \
    --defaultSchemaName=$defaultSchemaName \
    update

echo "Starting server..."
echo "Project: ${PROJECT_NAME}"

cd /backend_project
# Run the built executable
./${PROJECT_NAME}
