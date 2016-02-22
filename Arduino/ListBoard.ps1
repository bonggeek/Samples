
$f = Get-ChildItem -Path $args[0] -Filter "boards.txt" -Recurse
foreach($file in $f)
{
    Write-Host "For file" $file.FullName
    
    foreach ($l in get-content $file.FullName) {
        if($l.Contains(".name")) {
            $b = $l.Split('=')[1];
        }

        if($l.Contains(".board")) {
            $s = [string]::Format("{0,-40}ARDUINO_{1}", $b, ($l.Split('=')[1]));
            Write-Host $s
        }

    }
}