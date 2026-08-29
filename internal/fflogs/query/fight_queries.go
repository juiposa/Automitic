package query

const GetFightPlayersQuery = `
query($reportId: String, $fightNumber: Int) {
    reportData {
        report(code: $reportId) {
            masterData {
				playerDetails(fightIDs: [$fightNumber])
            }
        }
    }
}
`

const GetPaginatedFightEvents = `
query($reportId: String, $fightNumber: Int, $startTime: Float, $endTime: Float) {
    reportData {
        report(code: $reportId) {
            events (limit:10000, fightIDs:[$fightNumber], startTime: $startTime, endTime: $endTime) {
                nextPageTimestamp
                data
            }
        }
    }
}
`
