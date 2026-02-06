package main

import (
	"context"
	"log"

	sdk "github.com/apito-io/go-apito-plugin-sdk"
)

func main() {
	log.Printf("🎯 [hc-storage-plugin] Starting File Storage plugin...")

	plugin := sdk.Init("hc-storage-plugin", "1.0.0", "apito-plugin-key")

	statusType := sdk.NewObjectType("StorageStatus", "Storage plugin status").
		AddStringField("status", "Plugin status", false).
		AddStringField("version", "Plugin version", false).
		Build()

	plugin.RegisterQuery("getStorageStatus",
		sdk.ComplexObjectField("Get storage plugin status", statusType),
		func(ctx context.Context, rawArgs map[string]interface{}) (interface{}, error) {
			return map[string]interface{}{"status": "ready", "version": "1.0.0"}, nil
		})

	log.Printf("🚀 [hc-storage-plugin] Plugin ready")
	plugin.Serve()
}
