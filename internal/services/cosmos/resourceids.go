// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cosmos

//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=CassandraKeyspace -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.DocumentDB/databaseAccounts/acc1/cassandraKeyspaces/keyspace1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=CassandraTable -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.DocumentDB/databaseAccounts/acc1/cassandraKeyspaces/keyspace1/tables/table1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=DatabaseAccount -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.DocumentDB/databaseAccounts/acc1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=GremlinDatabase -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.DocumentDB/databaseAccounts/acc1/gremlinDatabases/database1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=GremlinGraph -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.DocumentDB/databaseAccounts/acc1/gremlinDatabases/database1/graphs/graph1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=MongodbCollection -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.DocumentDB/databaseAccounts/acc1/mongodbDatabases/db1/collections/coll1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=MongodbDatabase -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.DocumentDB/databaseAccounts/acc1/mongodbDatabases/db1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=NotebookWorkspace -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.DocumentDB/databaseAccounts/account1/notebookWorkspaces/notebookWorkspace1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=RestorableDatabaseAccount -id=/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.DocumentDB/locations/location1/restorableDatabaseAccounts/account1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=SqlContainer -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.DocumentDB/databaseAccounts/acc1/sqlDatabases/db1/containers/container1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=SqlDatabase -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.DocumentDB/databaseAccounts/acc1/sqlDatabases/db1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=SqlFunction -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.DocumentDB/databaseAccounts/account1/sqlDatabases/database1/containers/container1/userDefinedFunctions/userDefinedFunction1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=SqlRoleAssignment -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.DocumentDB/databaseAccounts/account1/sqlRoleAssignments/roleAssignment1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=SqlRoleDefinition -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.DocumentDB/databaseAccounts/account1/sqlRoleDefinitions/def1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=SqlStoredProcedure -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.DocumentDB/databaseAccounts/acc1/sqlDatabases/db1/containers/container1/storedProcedures/sproc1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=SqlTrigger -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.DocumentDB/databaseAccounts/account1/sqlDatabases/database1/containers/container1/triggers/trigger1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=Table -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.DocumentDB/databaseAccounts/acc1/tables/table1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=CassandraCluster -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.DocumentDB/cassandraClusters/cluster1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=CassandraDatacenter -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.DocumentDB/cassandraClusters/cluster1/dataCenters/dc1
