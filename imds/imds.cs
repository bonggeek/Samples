ing System;

using System.IO;
using System.Net;
using Newtonsoft.Json;

namespace imds
{
    class Program
    {
        static void Main(string[] args)
        {
            HttpWebRequest request = WebRequest.Create("http://169.254.169.254/metadata/instance?api-version=2017-04-02") as HttpWebRequest;
            request.Headers.Add("Metadata: True");
            // Get response
            using (HttpWebResponse response = request.GetResponse() as HttpWebResponse)
            {
                // Get the response stream
                StreamReader reader = new StreamReader(response.GetResponseStream());
                // Console application output
                string responseStr = reader.ReadToEnd();   
                // pretty print string           
                string json = JsonConvert.SerializeObject(JsonConvert.DeserializeObject(responseStr), Formatting.Indented);
                Console.WriteLine(json);
            }
        }
    }
}

