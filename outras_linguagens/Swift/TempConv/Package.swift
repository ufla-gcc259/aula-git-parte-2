// swift-tools-version: 5.6
// The swift-tools-version declares the minimum version of Swift required to build this package.

import PackageDescription

let package: Package = Package(
    name: "TempConv",
    products: [
        .library(
            name: "TempConv",
            targets: ["TempConv"]
        ),
        .executable(
            name: "SampleCLIApp",
            targets: ["SampleCLIApp"]
        )
    ],
    dependencies: [],
    targets: [
        .target(
            name: "TempConv",
            dependencies: []
        ),
        .executableTarget(
            name: "SampleCLIApp",
            dependencies: ["TempConv"]
        )
    ]
)
