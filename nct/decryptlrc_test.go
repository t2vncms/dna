package nct

import (
	"dna"
)

func ExampleDecryptLRC() {
	var data dna.String = "0BAF2564A1C1AF9E3AC9A5F6FDC42CA0EBA4A3584FD54C91CF3983538CC761B0EB54E77BFF3994C6188C1C26FEB004C745119E11A53E618DA62754364E842D97900E0F531E872BA76652FAE77EB2FD764DBE11966FA1BF4A96C86BC28F33754B60BF102E9CEF0578531B41315B6398F8A6B35234881FEFC172134FF8B2832D6BEE9A757846AB4020B32E8505730204154480DD0938EE69427D2E42EB48787E9BAD63FAEC9EFA15E0619B6EE55993B824A85F9D2E1E957BCB8EA6B9C1AF45B6C0F607D6D32FFC7A12316829F8B063F86B52C168A2830AECE58E518D8D43D5E9FCE9F771CF716D9741926199613733D2BDEB0FA696FD00910A422DFCCC80E376454DCDEE45439C9F7C52A9453354720BAC5FE03B74C70F4EF98900B82E3A66B4A25F822417107B84F0A77A6DB454896D4F5C969DFF721A494BE8576A7716BE23B3B40203DD07B29DB241F808B2D55629F3C31F13FE780A93F333DA92E227BC6DAEA733F25B8EAF2ACEE2DA5C2F7526103F515B515BDDC57827E45D330096DC1036DA68D723D879A42D5B35CEA9E88CD819121ED9D01D68FFBD40AF2D36073125A344B66B4476F53A91B5E0E2F10E29B44AB1CF0E5D8077DE71FD85B4787EE4DF3AA18C03323A9C4AF8B2E75E049F4FDD470B56C60FCF029422C4D0A12995CC367238E4A6544DF27607CBB0EB8394DC7D320F074F8B323D961EB152999E8E680BBF0B8B2CA5AC435907CAA9B4D3BC1A105B7C02D3CC9EBA76BAA1F33DCBD19D087D7D49B177239991C54E7E600D6CA23298BDA9B32DEACB370BB05192D68F8A372AC376337017AC65B08CBAB24A75615B89E9D23C24B36A47312FA7915E842EAE7BB89C8A7D60ED8A8208B2B53A41664DBA6D0F66511F3BB80A21E0F812554F3FAC8D966332B6B76FA52E3FD3F065D8DD2C287FDEC06C5FCAD9DCC5E9319D3D119A1E23F45077080576D09F1233E7C1AB8C9DA3C898CA5BCBA2443C8077BFAEE2816C3BF54B4A3F51AADDA5BF72DF791C508761411210DA2E1513F3B36393E62692179DF39D563EDED84002EFE51E7DB3F1006473ECE8BE34B86D5FB313AB1E84E3D40E4C9328D7EBF6A2961EC56642222172708DB17006AA31C30C3E206FB2AFA3CEDFA615FF188A581844930C40AD3DDFACA2C238FEEA2349191AF8379F0BAFA487ED7CDBBD38726E9EC6CE137B33FC5C8620F053134291EA9C1AD4A447C2B8D00FB13C05AFC92746F88EFD2E35B39A18F1B3586A6AD1DC12B1DCA0D5B553CD4CC31E63340DD769DC8C1F40A785D42F16C36E1C29DC4EB464770FCA54586BD4563B83E15381167DECB63962359CCC836FBB367C78186D45BDDA8254952C76C30EDEBA5EFA0DB538C4969CF6EE4EAD7FC6B08D1B8E6DA1BCFFD2B38806926663A24B9CB004C1D3695B7072A4C95AEC078AF6FCC292F53B860A4E3E86656833B75F8661F381742A3D628F8B444257E81609566D76D6BBAC87817530065F8222237DAB1C0B39D4080177D558ABBCC991A0519A9703F2A43486EECEBC9BE02578287995A9183C04E353E5DEBB746E2EAAC8316DF517B00E8CFDB73F947F2A9B339915FA6D58EFFAAB9A219C57C20D696E582E5D084DAFA0A250B8A2B0CE9FF9D0D170399EA553D4A6E10A7F4653C030909E2F0261CDA09893E567682069F1C58EF848DC17B583B0E2522730944A931697E4FEEA8313E5B2D8DF708936EBC59677A446515597772076CF11AC3C86A540CFDEE5F425972E1CE8F71083C64F2B046F03BC44E9ABECFFF58223C61A43C3A6D011FE06CE82BF431ECC3ADBA29F747E6814400DDA77BCC4C1EB9AD60B395EBE7763CEC569B6EC6CF0958237CB0F61F5FB838D553CE66CADEF376EF4D63C52AF894D05CD283BCE855AF9D4F7B435B45238516C277C844376A5AE30F2F4486099DEFB1958954E4884ED1DEB5CDA6F87FCB3DC60042770A12ADB100AF2E88A32AB7C13F46B57C78C06D32FAFD449D1A131F01C8A712AAF93D78C2F35EA77AF260454D071A0A0C64B06EA8C813E73298EE9FDF0A1A7CBB7291841B01F1CC12EE0C5978FC7B51C4AAD36B169F8193ACAF09F2983296E3DF6BB590BF1D70D98006B9BEDA0705752AC0C9C000B5A8E9560E4A460C15A1096DDFEA465BC0D8BD4E28A01A41437DC987243EF973643D51AAA323AC6258E99666786EF58DB080CA6ED888E31D9D53504C6886977F9BFB12C82CF2A44366086D1B90F427BA80ED7350FD4D503CFBE03099F3509BB55FE0A257BD93FC868C29F1839609B1408B1C0A03B533A591843B7B08317AE7DC048147B950A885EF72EA39B3455073AE90E8C4DBCED8E17A0D28D7261AC5A715F3D4A6836BFB298A9E817F25169DF2A9157D541D201A115BFDB08E69D92FBFA6D2AAFF8F21BAE49142E63450BF4472645394D8D77DF7A4DB447D3A4D624F6647E0EDCB23C8F121720B9C22FD04DE869B65B6D44C98E89D86A1D1A150EAAF40E6811E53849E5FD8D6A599C9D12EC701FBE1280F9A8AA8E55744A96899B234547F04D59D4B23C9203073F6DE922A57A9D6BE70CBDAAC8D5B31637BB9B58EC87B22ACAFE390C48C853457B97E9E810C2AC03EBB74D9F7B660E4BA608ADD89FE6AC974C371054716B7068D2DA576A5436CA16B291B76F00F4A35CDBE58BB04387062930E89ED31F419337E8549E1DEE2BA17D4AA9B0206AEDA4017C47987E812853ED9BEC37C2ACA941147AF0627BDB893D90D8CB7CA1D363A2087D0652B70BCD4407A706C55FEBB77B916996023E04253F796F0818C4338B90E6471C912D22CA1FCFDAB2ADB0964EC9EB9D853DED8D1F62E0CBD19E1851E6BEB48D527174F6C3E3BD88E5E3FC11A440CBAEFFB7845F08BA75F0E08F0670AD934FDA161470B49A44E3ABDC709534BD28C0E02225B60CEE57151B5240608068F7E1C4314394B81C941904371E58038D3FDCA479A3857A6778C0D75C51E8625682AD9E25960E0073F4E96CCB8A6759C247A537B08A16BA49ABAC0793603D1CD4883C9DE90B2A3B5C998453868DE295B16FF22FB6F1D56222339F11AC6FAEF006507035B8AECAF1DB546307B63040"

	lrc, err := DecryptLRC(data)
	var expectedLrc dna.String = `[00:02.06]Ca Khúc : La La la
[00:06.06]Trình Bày : Ông Cao Thắng
[00:10.06]Lyris By Hải Biển
[00:14.06]
[00:19.06]Tôi cô đơn 
[00:19.86]đến đây cũng lâu rồi
[00:22.84]chỉ tại cuộc tình 
[00:24.39]vừa trãi qua thôi
[00:26.61]qua bao lâu 
[00:27.89]tôi lại đến đây rồi
[00:30.76]và hôm nay tôi 
[00:32.32]cần một người yêu tôi
[00:34.39]la la la la la la
[00:36.15]xem hôm nay 
[00:36.87]vận may thế nào
[00:38.14]la la la la la la
[00:40.20]thử lại một lần xem sao
[00:42.12]la la la la la la
[00:44.28]hãy đến cạnh 
[00:45.35]bên tôi nào
[00:46.17]la la la la la la
[00:48.15]mạnh dạn một lần xem sao
[02:25.71][01:29.88][00:49.78]Đừng ngại cứ bước 
[02:26.61][01:30.60][00:50.64]đến bên tôi 
[02:27.81][01:31.90][00:52.03]rồi cũng sẽ quen dần
[02:29.26][01:33.27][00:53.46]rồi cùng nhịp thở 
[02:30.46][01:34.59][00:54.82]thở ra lan tỏa 
[02:32.19][01:36.12][00:56.38]từ bản thân
[02:33.33][01:37.38][00:57.51]và từng giây phút nóng lên
[02:35.32][01:39.46][00:59.45] cho ta vực lên tinh thần
[02:37.17][01:41.19][01:01.27]bay lên bay lên 
[02:38.79][01:42.71][01:02.90]cho ta bay xa 
[02:57.11][02:48.77][02:42.02]
[01:45.44][02:58.00][01:45.21][01:05.26]everybody say 
[03:09.16][03:07.19][03:05.32][03:03.01][03:01.00][02:59.26][02:58.31][03:11.60]everybody say  la la 
[02:46.59][02:44.58][02:42.17][01:50.59][01:48.54][01:46.33][01:11.14][01:08.52][01:06.44]la la la la la la
[01:53.08][01:13.10]everybody say 
[02:54.67][02:52.62][02:50.35][01:59.38][01:56.48][01:54.35][01:14.28]la la la la la la
[02:02.21]
[02:10.08]
[02:11.15]Không gian quanh đây
[02:12.46]đã tan dần dần
[02:14.91]bỗng thấy nôn nao
[02:16.31]buộc lòng đôi chân
[02:18.83]chim nào bay cao 
[02:20.46]bay cao vượt rào
[02:23.07]bỗng thấy nôn nao nôn nao
[01:16.21]xem hôm nay 
[01:16.90]vận may thế nào
[01:18.10]la la la la la la
[01:20.06]thử lại một lần xem sao
[01:22.14]la la la la la la
[01:24.09]ai muốn cạnh bên tôi nào
[01:26.14]la la la la la la
[01:28.16]mạnh dạn một lần xem sao`
	if err == nil {
		if lrc != expectedLrc {
			panic("Wrong")
		}
	}

}
