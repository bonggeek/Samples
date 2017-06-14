Imports System.IO
Imports System.Net
Imports Newtonsoft.Json

Module Module1

    Sub Main()
        Dim request As HttpWebRequest = WebRequest.Create("http://169.254.169.254/metadata/instance?api-version=2017-04-02")
        request.Headers.Add("Metadata: True")
        Dim response As HttpWebResponse = request.GetResponse()
        Dim dataStream As Stream = response.GetResponseStream()
        Dim reader As New StreamReader(dataStream)
        Dim responseFromServer As String = reader.ReadToEnd()

        Dim json As String = JsonConvert.SerializeObject(JsonConvert.DeserializeObject(responseFromServer), Formatting.Indented)
        Console.WriteLine(json)
        reader.Close()
        dataStream.Close()
        response.Close()
    End Sub

End Module
        
