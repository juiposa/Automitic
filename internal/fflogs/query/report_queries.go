package query

const GetFightsQuery = `
query($reportId: String) {
    reportData {
        report(code: $reportId) {
            fights {
                id
                startTime
                endTime
                friendlyPlayers
            }
        }
    }
}
`

const GetMetadataQuery = `
query($reportId: String) {
    reportData {
        report(code: $reportId) {
            masterData {
                actors {
                    id
                    gameID
                    name
                    server
                    type
                    subType
                }
            }
        }
    }
}
`
