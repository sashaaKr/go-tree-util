Simple util to show folder structure tree:

```
├───main.go (1881b)        
├───main_test.go (1318b)    
└───testdata                
	├───project               
	│	├───file.txt (19b)      
	│	└───gopher.png (70372b) 
	├───static                
	│	├───css                 
	│	│	└───body.css (28b)    
	│	├───html                
	│	│	└───index.html (57b)  
	│	└───js                  
	│		└───site.js (10b)     
	├───zline                 
	│	└───empty.txt (empty)   
	└───zzfile.txt (empty) 
```
### Run program:
```
go run main.go <pathToRoot> [-f]
```
### Run tests:
```
go test -v
```

### Arguments
`pathToRoot` - start point of scan

`-f` - show files
