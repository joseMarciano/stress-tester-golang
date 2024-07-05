package application

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"net/http"
	"os"
	"stress-test/internal/model"
)

type ReportPrinterService struct {
}

func (s ReportPrinterService) Print(report *model.Report) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Total Requests", "Total Requests 200", "Total Spent Time(ms)"})
	t.AppendRows([]table.Row{
		{report.TotalRequests, report.TotalRequestByStatusCode(http.StatusOK), report.TotalSpentTime},
	})
	t.AppendFooter(table.Row{""})
	t.AppendSeparator()
	t.SetStyle(table.StyleBold)
	t.Render()

	t.ResetHeaders()
	t.ResetRows()
	t.ResetFooters()
	t.AppendHeader(table.Row{"Status Code", "Total Requests"})
	for key, value := range report.RequestMap {
		t.AppendRow(table.Row{key, value})
	}

	t.Render()
}
