package engine

// **************************************** 
// Author(s): Joo-Von Kim, C2N, CNRS/Univ. Paris-Saclay
//
// This module projects the magnetization onto user-supplied transverse directions,
// delta_mx, delta_my (obtained, e.g., from the relaxed magnetization state), with
// a spatial convolution with the user-supplied masks psi_kn. It returns the amplitudes
// 
// 	a_kxn = int_dV {psi_kn (m . delta_mx)}
//	a_kyn = int_dV {psi_kn (m . delta_my)}
//
// with up to 500 modes (n = 000 to 499).
//
// The user-supplied masks/vector fields can be added in the source .mx3 file with
//	psi_kn.Add( LoadFile(("psi_kn_file.ovf"),1) )	
//	delta_mx.Add( LoadFile("delta_mx_file.ovf"), 1 )
//	etc.
//
// Acknowledgements:
// This work was supported by Horizon Europe Research and Innovation Programme of the
// European Commission under grant agreement No. 101070290 (NIMFEIA). 
//
// **************************************** 


import (
	"github.com/mumax/3/cuda"
)


var (
//  Transverse magnetization is already defined in ext_eigenmodeprojection.go
//	delta_mx	= NewExcitation("delta_mx", "", "Transverse magnetization 1")
//	delta_my	= NewExcitation("delta_my", "", "Transverse magnetization 2")

	psi_k000		= NewScalarExcitation("psi_k000", "", "Eigenmode spatial profile")
	psi_k001		= NewScalarExcitation("psi_k001", "", "Eigenmode spatial profile")
	psi_k002		= NewScalarExcitation("psi_k002", "", "Eigenmode spatial profile")
	psi_k003		= NewScalarExcitation("psi_k003", "", "Eigenmode spatial profile")
	psi_k004		= NewScalarExcitation("psi_k004", "", "Eigenmode spatial profile")
	psi_k005		= NewScalarExcitation("psi_k005", "", "Eigenmode spatial profile")
	psi_k006		= NewScalarExcitation("psi_k006", "", "Eigenmode spatial profile")
	psi_k007		= NewScalarExcitation("psi_k007", "", "Eigenmode spatial profile")
	psi_k008		= NewScalarExcitation("psi_k008", "", "Eigenmode spatial profile")
	psi_k009		= NewScalarExcitation("psi_k009", "", "Eigenmode spatial profile")
	psi_k010		= NewScalarExcitation("psi_k010", "", "Eigenmode spatial profile")
	psi_k011		= NewScalarExcitation("psi_k011", "", "Eigenmode spatial profile")
	psi_k012		= NewScalarExcitation("psi_k012", "", "Eigenmode spatial profile")
	psi_k013		= NewScalarExcitation("psi_k013", "", "Eigenmode spatial profile")
	psi_k014		= NewScalarExcitation("psi_k014", "", "Eigenmode spatial profile")
	psi_k015		= NewScalarExcitation("psi_k015", "", "Eigenmode spatial profile")
	psi_k016		= NewScalarExcitation("psi_k016", "", "Eigenmode spatial profile")
	psi_k017		= NewScalarExcitation("psi_k017", "", "Eigenmode spatial profile")
	psi_k018		= NewScalarExcitation("psi_k018", "", "Eigenmode spatial profile")
	psi_k019		= NewScalarExcitation("psi_k019", "", "Eigenmode spatial profile")
	psi_k020		= NewScalarExcitation("psi_k020", "", "Eigenmode spatial profile")
	psi_k021		= NewScalarExcitation("psi_k021", "", "Eigenmode spatial profile")
	psi_k022		= NewScalarExcitation("psi_k022", "", "Eigenmode spatial profile")
	psi_k023		= NewScalarExcitation("psi_k023", "", "Eigenmode spatial profile")
	psi_k024		= NewScalarExcitation("psi_k024", "", "Eigenmode spatial profile")
	psi_k025		= NewScalarExcitation("psi_k025", "", "Eigenmode spatial profile")
	psi_k026		= NewScalarExcitation("psi_k026", "", "Eigenmode spatial profile")
	psi_k027		= NewScalarExcitation("psi_k027", "", "Eigenmode spatial profile")
	psi_k028		= NewScalarExcitation("psi_k028", "", "Eigenmode spatial profile")
	psi_k029		= NewScalarExcitation("psi_k029", "", "Eigenmode spatial profile")
	psi_k030		= NewScalarExcitation("psi_k030", "", "Eigenmode spatial profile")
	psi_k031		= NewScalarExcitation("psi_k031", "", "Eigenmode spatial profile")
	psi_k032		= NewScalarExcitation("psi_k032", "", "Eigenmode spatial profile")
	psi_k033		= NewScalarExcitation("psi_k033", "", "Eigenmode spatial profile")
	psi_k034		= NewScalarExcitation("psi_k034", "", "Eigenmode spatial profile")
	psi_k035		= NewScalarExcitation("psi_k035", "", "Eigenmode spatial profile")
	psi_k036		= NewScalarExcitation("psi_k036", "", "Eigenmode spatial profile")
	psi_k037		= NewScalarExcitation("psi_k037", "", "Eigenmode spatial profile")
	psi_k038		= NewScalarExcitation("psi_k038", "", "Eigenmode spatial profile")
	psi_k039		= NewScalarExcitation("psi_k039", "", "Eigenmode spatial profile")
	psi_k040		= NewScalarExcitation("psi_k040", "", "Eigenmode spatial profile")
	psi_k041		= NewScalarExcitation("psi_k041", "", "Eigenmode spatial profile")
	psi_k042		= NewScalarExcitation("psi_k042", "", "Eigenmode spatial profile")
	psi_k043		= NewScalarExcitation("psi_k043", "", "Eigenmode spatial profile")
	psi_k044		= NewScalarExcitation("psi_k044", "", "Eigenmode spatial profile")
	psi_k045		= NewScalarExcitation("psi_k045", "", "Eigenmode spatial profile")
	psi_k046		= NewScalarExcitation("psi_k046", "", "Eigenmode spatial profile")
	psi_k047		= NewScalarExcitation("psi_k047", "", "Eigenmode spatial profile")
	psi_k048		= NewScalarExcitation("psi_k048", "", "Eigenmode spatial profile")
	psi_k049		= NewScalarExcitation("psi_k049", "", "Eigenmode spatial profile")
	psi_k050		= NewScalarExcitation("psi_k050", "", "Eigenmode spatial profile")
	psi_k051		= NewScalarExcitation("psi_k051", "", "Eigenmode spatial profile")
	psi_k052		= NewScalarExcitation("psi_k052", "", "Eigenmode spatial profile")
	psi_k053		= NewScalarExcitation("psi_k053", "", "Eigenmode spatial profile")
	psi_k054		= NewScalarExcitation("psi_k054", "", "Eigenmode spatial profile")
	psi_k055		= NewScalarExcitation("psi_k055", "", "Eigenmode spatial profile")
	psi_k056		= NewScalarExcitation("psi_k056", "", "Eigenmode spatial profile")
	psi_k057		= NewScalarExcitation("psi_k057", "", "Eigenmode spatial profile")
	psi_k058		= NewScalarExcitation("psi_k058", "", "Eigenmode spatial profile")
	psi_k059		= NewScalarExcitation("psi_k059", "", "Eigenmode spatial profile")
	psi_k060		= NewScalarExcitation("psi_k060", "", "Eigenmode spatial profile")
	psi_k061		= NewScalarExcitation("psi_k061", "", "Eigenmode spatial profile")
	psi_k062		= NewScalarExcitation("psi_k062", "", "Eigenmode spatial profile")
	psi_k063		= NewScalarExcitation("psi_k063", "", "Eigenmode spatial profile")
	psi_k064		= NewScalarExcitation("psi_k064", "", "Eigenmode spatial profile")
	psi_k065		= NewScalarExcitation("psi_k065", "", "Eigenmode spatial profile")
	psi_k066		= NewScalarExcitation("psi_k066", "", "Eigenmode spatial profile")
	psi_k067		= NewScalarExcitation("psi_k067", "", "Eigenmode spatial profile")
	psi_k068		= NewScalarExcitation("psi_k068", "", "Eigenmode spatial profile")
	psi_k069		= NewScalarExcitation("psi_k069", "", "Eigenmode spatial profile")
	psi_k070		= NewScalarExcitation("psi_k070", "", "Eigenmode spatial profile")
	psi_k071		= NewScalarExcitation("psi_k071", "", "Eigenmode spatial profile")
	psi_k072		= NewScalarExcitation("psi_k072", "", "Eigenmode spatial profile")
	psi_k073		= NewScalarExcitation("psi_k073", "", "Eigenmode spatial profile")
	psi_k074		= NewScalarExcitation("psi_k074", "", "Eigenmode spatial profile")
	psi_k075		= NewScalarExcitation("psi_k075", "", "Eigenmode spatial profile")
	psi_k076		= NewScalarExcitation("psi_k076", "", "Eigenmode spatial profile")
	psi_k077		= NewScalarExcitation("psi_k077", "", "Eigenmode spatial profile")
	psi_k078		= NewScalarExcitation("psi_k078", "", "Eigenmode spatial profile")
	psi_k079		= NewScalarExcitation("psi_k079", "", "Eigenmode spatial profile")
	psi_k080		= NewScalarExcitation("psi_k080", "", "Eigenmode spatial profile")
	psi_k081		= NewScalarExcitation("psi_k081", "", "Eigenmode spatial profile")
	psi_k082		= NewScalarExcitation("psi_k082", "", "Eigenmode spatial profile")
	psi_k083		= NewScalarExcitation("psi_k083", "", "Eigenmode spatial profile")
	psi_k084		= NewScalarExcitation("psi_k084", "", "Eigenmode spatial profile")
	psi_k085		= NewScalarExcitation("psi_k085", "", "Eigenmode spatial profile")
	psi_k086		= NewScalarExcitation("psi_k086", "", "Eigenmode spatial profile")
	psi_k087		= NewScalarExcitation("psi_k087", "", "Eigenmode spatial profile")
	psi_k088		= NewScalarExcitation("psi_k088", "", "Eigenmode spatial profile")
	psi_k089		= NewScalarExcitation("psi_k089", "", "Eigenmode spatial profile")
	psi_k090		= NewScalarExcitation("psi_k090", "", "Eigenmode spatial profile")
	psi_k091		= NewScalarExcitation("psi_k091", "", "Eigenmode spatial profile")
	psi_k092		= NewScalarExcitation("psi_k092", "", "Eigenmode spatial profile")
	psi_k093		= NewScalarExcitation("psi_k093", "", "Eigenmode spatial profile")
	psi_k094		= NewScalarExcitation("psi_k094", "", "Eigenmode spatial profile")
	psi_k095		= NewScalarExcitation("psi_k095", "", "Eigenmode spatial profile")
	psi_k096		= NewScalarExcitation("psi_k096", "", "Eigenmode spatial profile")
	psi_k097		= NewScalarExcitation("psi_k097", "", "Eigenmode spatial profile")
	psi_k098		= NewScalarExcitation("psi_k098", "", "Eigenmode spatial profile")
	psi_k099		= NewScalarExcitation("psi_k099", "", "Eigenmode spatial profile")
	psi_k100		= NewScalarExcitation("psi_k100", "", "Eigenmode spatial profile")
	psi_k101		= NewScalarExcitation("psi_k101", "", "Eigenmode spatial profile")
	psi_k102		= NewScalarExcitation("psi_k102", "", "Eigenmode spatial profile")
	psi_k103		= NewScalarExcitation("psi_k103", "", "Eigenmode spatial profile")
	psi_k104		= NewScalarExcitation("psi_k104", "", "Eigenmode spatial profile")
	psi_k105		= NewScalarExcitation("psi_k105", "", "Eigenmode spatial profile")
	psi_k106		= NewScalarExcitation("psi_k106", "", "Eigenmode spatial profile")
	psi_k107		= NewScalarExcitation("psi_k107", "", "Eigenmode spatial profile")
	psi_k108		= NewScalarExcitation("psi_k108", "", "Eigenmode spatial profile")
	psi_k109		= NewScalarExcitation("psi_k109", "", "Eigenmode spatial profile")
	psi_k110		= NewScalarExcitation("psi_k110", "", "Eigenmode spatial profile")
	psi_k111		= NewScalarExcitation("psi_k111", "", "Eigenmode spatial profile")
	psi_k112		= NewScalarExcitation("psi_k112", "", "Eigenmode spatial profile")
	psi_k113		= NewScalarExcitation("psi_k113", "", "Eigenmode spatial profile")
	psi_k114		= NewScalarExcitation("psi_k114", "", "Eigenmode spatial profile")
	psi_k115		= NewScalarExcitation("psi_k115", "", "Eigenmode spatial profile")
	psi_k116		= NewScalarExcitation("psi_k116", "", "Eigenmode spatial profile")
	psi_k117		= NewScalarExcitation("psi_k117", "", "Eigenmode spatial profile")
	psi_k118		= NewScalarExcitation("psi_k118", "", "Eigenmode spatial profile")
	psi_k119		= NewScalarExcitation("psi_k119", "", "Eigenmode spatial profile")
	psi_k120		= NewScalarExcitation("psi_k120", "", "Eigenmode spatial profile")
	psi_k121		= NewScalarExcitation("psi_k121", "", "Eigenmode spatial profile")
	psi_k122		= NewScalarExcitation("psi_k122", "", "Eigenmode spatial profile")
	psi_k123		= NewScalarExcitation("psi_k123", "", "Eigenmode spatial profile")
	psi_k124		= NewScalarExcitation("psi_k124", "", "Eigenmode spatial profile")
	psi_k125		= NewScalarExcitation("psi_k125", "", "Eigenmode spatial profile")
	psi_k126		= NewScalarExcitation("psi_k126", "", "Eigenmode spatial profile")
	psi_k127		= NewScalarExcitation("psi_k127", "", "Eigenmode spatial profile")
	psi_k128		= NewScalarExcitation("psi_k128", "", "Eigenmode spatial profile")
	psi_k129		= NewScalarExcitation("psi_k129", "", "Eigenmode spatial profile")
	psi_k130		= NewScalarExcitation("psi_k130", "", "Eigenmode spatial profile")
	psi_k131		= NewScalarExcitation("psi_k131", "", "Eigenmode spatial profile")
	psi_k132		= NewScalarExcitation("psi_k132", "", "Eigenmode spatial profile")
	psi_k133		= NewScalarExcitation("psi_k133", "", "Eigenmode spatial profile")
	psi_k134		= NewScalarExcitation("psi_k134", "", "Eigenmode spatial profile")
	psi_k135		= NewScalarExcitation("psi_k135", "", "Eigenmode spatial profile")
	psi_k136		= NewScalarExcitation("psi_k136", "", "Eigenmode spatial profile")
	psi_k137		= NewScalarExcitation("psi_k137", "", "Eigenmode spatial profile")
	psi_k138		= NewScalarExcitation("psi_k138", "", "Eigenmode spatial profile")
	psi_k139		= NewScalarExcitation("psi_k139", "", "Eigenmode spatial profile")
	psi_k140		= NewScalarExcitation("psi_k140", "", "Eigenmode spatial profile")
	psi_k141		= NewScalarExcitation("psi_k141", "", "Eigenmode spatial profile")
	psi_k142		= NewScalarExcitation("psi_k142", "", "Eigenmode spatial profile")
	psi_k143		= NewScalarExcitation("psi_k143", "", "Eigenmode spatial profile")
	psi_k144		= NewScalarExcitation("psi_k144", "", "Eigenmode spatial profile")
	psi_k145		= NewScalarExcitation("psi_k145", "", "Eigenmode spatial profile")
	psi_k146		= NewScalarExcitation("psi_k146", "", "Eigenmode spatial profile")
	psi_k147		= NewScalarExcitation("psi_k147", "", "Eigenmode spatial profile")
	psi_k148		= NewScalarExcitation("psi_k148", "", "Eigenmode spatial profile")
	psi_k149		= NewScalarExcitation("psi_k149", "", "Eigenmode spatial profile")
	psi_k150		= NewScalarExcitation("psi_k150", "", "Eigenmode spatial profile")
	psi_k151		= NewScalarExcitation("psi_k151", "", "Eigenmode spatial profile")
	psi_k152		= NewScalarExcitation("psi_k152", "", "Eigenmode spatial profile")
	psi_k153		= NewScalarExcitation("psi_k153", "", "Eigenmode spatial profile")
	psi_k154		= NewScalarExcitation("psi_k154", "", "Eigenmode spatial profile")
	psi_k155		= NewScalarExcitation("psi_k155", "", "Eigenmode spatial profile")
	psi_k156		= NewScalarExcitation("psi_k156", "", "Eigenmode spatial profile")
	psi_k157		= NewScalarExcitation("psi_k157", "", "Eigenmode spatial profile")
	psi_k158		= NewScalarExcitation("psi_k158", "", "Eigenmode spatial profile")
	psi_k159		= NewScalarExcitation("psi_k159", "", "Eigenmode spatial profile")
	psi_k160		= NewScalarExcitation("psi_k160", "", "Eigenmode spatial profile")
	psi_k161		= NewScalarExcitation("psi_k161", "", "Eigenmode spatial profile")
	psi_k162		= NewScalarExcitation("psi_k162", "", "Eigenmode spatial profile")
	psi_k163		= NewScalarExcitation("psi_k163", "", "Eigenmode spatial profile")
	psi_k164		= NewScalarExcitation("psi_k164", "", "Eigenmode spatial profile")
	psi_k165		= NewScalarExcitation("psi_k165", "", "Eigenmode spatial profile")
	psi_k166		= NewScalarExcitation("psi_k166", "", "Eigenmode spatial profile")
	psi_k167		= NewScalarExcitation("psi_k167", "", "Eigenmode spatial profile")
	psi_k168		= NewScalarExcitation("psi_k168", "", "Eigenmode spatial profile")
	psi_k169		= NewScalarExcitation("psi_k169", "", "Eigenmode spatial profile")
	psi_k170		= NewScalarExcitation("psi_k170", "", "Eigenmode spatial profile")
	psi_k171		= NewScalarExcitation("psi_k171", "", "Eigenmode spatial profile")
	psi_k172		= NewScalarExcitation("psi_k172", "", "Eigenmode spatial profile")
	psi_k173		= NewScalarExcitation("psi_k173", "", "Eigenmode spatial profile")
	psi_k174		= NewScalarExcitation("psi_k174", "", "Eigenmode spatial profile")
	psi_k175		= NewScalarExcitation("psi_k175", "", "Eigenmode spatial profile")
	psi_k176		= NewScalarExcitation("psi_k176", "", "Eigenmode spatial profile")
	psi_k177		= NewScalarExcitation("psi_k177", "", "Eigenmode spatial profile")
	psi_k178		= NewScalarExcitation("psi_k178", "", "Eigenmode spatial profile")
	psi_k179		= NewScalarExcitation("psi_k179", "", "Eigenmode spatial profile")
	psi_k180		= NewScalarExcitation("psi_k180", "", "Eigenmode spatial profile")
	psi_k181		= NewScalarExcitation("psi_k181", "", "Eigenmode spatial profile")
	psi_k182		= NewScalarExcitation("psi_k182", "", "Eigenmode spatial profile")
	psi_k183		= NewScalarExcitation("psi_k183", "", "Eigenmode spatial profile")
	psi_k184		= NewScalarExcitation("psi_k184", "", "Eigenmode spatial profile")
	psi_k185		= NewScalarExcitation("psi_k185", "", "Eigenmode spatial profile")
	psi_k186		= NewScalarExcitation("psi_k186", "", "Eigenmode spatial profile")
	psi_k187		= NewScalarExcitation("psi_k187", "", "Eigenmode spatial profile")
	psi_k188		= NewScalarExcitation("psi_k188", "", "Eigenmode spatial profile")
	psi_k189		= NewScalarExcitation("psi_k189", "", "Eigenmode spatial profile")
	psi_k190		= NewScalarExcitation("psi_k190", "", "Eigenmode spatial profile")
	psi_k191		= NewScalarExcitation("psi_k191", "", "Eigenmode spatial profile")
	psi_k192		= NewScalarExcitation("psi_k192", "", "Eigenmode spatial profile")
	psi_k193		= NewScalarExcitation("psi_k193", "", "Eigenmode spatial profile")
	psi_k194		= NewScalarExcitation("psi_k194", "", "Eigenmode spatial profile")
	psi_k195		= NewScalarExcitation("psi_k195", "", "Eigenmode spatial profile")
	psi_k196		= NewScalarExcitation("psi_k196", "", "Eigenmode spatial profile")
	psi_k197		= NewScalarExcitation("psi_k197", "", "Eigenmode spatial profile")
	psi_k198		= NewScalarExcitation("psi_k198", "", "Eigenmode spatial profile")
	psi_k199		= NewScalarExcitation("psi_k199", "", "Eigenmode spatial profile")
	psi_k200		= NewScalarExcitation("psi_k200", "", "Eigenmode spatial profile")
	psi_k201		= NewScalarExcitation("psi_k201", "", "Eigenmode spatial profile")
	psi_k202		= NewScalarExcitation("psi_k202", "", "Eigenmode spatial profile")
	psi_k203		= NewScalarExcitation("psi_k203", "", "Eigenmode spatial profile")
	psi_k204		= NewScalarExcitation("psi_k204", "", "Eigenmode spatial profile")
	psi_k205		= NewScalarExcitation("psi_k205", "", "Eigenmode spatial profile")
	psi_k206		= NewScalarExcitation("psi_k206", "", "Eigenmode spatial profile")
	psi_k207		= NewScalarExcitation("psi_k207", "", "Eigenmode spatial profile")
	psi_k208		= NewScalarExcitation("psi_k208", "", "Eigenmode spatial profile")
	psi_k209		= NewScalarExcitation("psi_k209", "", "Eigenmode spatial profile")
	psi_k210		= NewScalarExcitation("psi_k210", "", "Eigenmode spatial profile")
	psi_k211		= NewScalarExcitation("psi_k211", "", "Eigenmode spatial profile")
	psi_k212		= NewScalarExcitation("psi_k212", "", "Eigenmode spatial profile")
	psi_k213		= NewScalarExcitation("psi_k213", "", "Eigenmode spatial profile")
	psi_k214		= NewScalarExcitation("psi_k214", "", "Eigenmode spatial profile")
	psi_k215		= NewScalarExcitation("psi_k215", "", "Eigenmode spatial profile")
	psi_k216		= NewScalarExcitation("psi_k216", "", "Eigenmode spatial profile")
	psi_k217		= NewScalarExcitation("psi_k217", "", "Eigenmode spatial profile")
	psi_k218		= NewScalarExcitation("psi_k218", "", "Eigenmode spatial profile")
	psi_k219		= NewScalarExcitation("psi_k219", "", "Eigenmode spatial profile")
	psi_k220		= NewScalarExcitation("psi_k220", "", "Eigenmode spatial profile")
	psi_k221		= NewScalarExcitation("psi_k221", "", "Eigenmode spatial profile")
	psi_k222		= NewScalarExcitation("psi_k222", "", "Eigenmode spatial profile")
	psi_k223		= NewScalarExcitation("psi_k223", "", "Eigenmode spatial profile")
	psi_k224		= NewScalarExcitation("psi_k224", "", "Eigenmode spatial profile")
	psi_k225		= NewScalarExcitation("psi_k225", "", "Eigenmode spatial profile")
	psi_k226		= NewScalarExcitation("psi_k226", "", "Eigenmode spatial profile")
	psi_k227		= NewScalarExcitation("psi_k227", "", "Eigenmode spatial profile")
	psi_k228		= NewScalarExcitation("psi_k228", "", "Eigenmode spatial profile")
	psi_k229		= NewScalarExcitation("psi_k229", "", "Eigenmode spatial profile")
	psi_k230		= NewScalarExcitation("psi_k230", "", "Eigenmode spatial profile")
	psi_k231		= NewScalarExcitation("psi_k231", "", "Eigenmode spatial profile")
	psi_k232		= NewScalarExcitation("psi_k232", "", "Eigenmode spatial profile")
	psi_k233		= NewScalarExcitation("psi_k233", "", "Eigenmode spatial profile")
	psi_k234		= NewScalarExcitation("psi_k234", "", "Eigenmode spatial profile")
	psi_k235		= NewScalarExcitation("psi_k235", "", "Eigenmode spatial profile")
	psi_k236		= NewScalarExcitation("psi_k236", "", "Eigenmode spatial profile")
	psi_k237		= NewScalarExcitation("psi_k237", "", "Eigenmode spatial profile")
	psi_k238		= NewScalarExcitation("psi_k238", "", "Eigenmode spatial profile")
	psi_k239		= NewScalarExcitation("psi_k239", "", "Eigenmode spatial profile")
	psi_k240		= NewScalarExcitation("psi_k240", "", "Eigenmode spatial profile")
	psi_k241		= NewScalarExcitation("psi_k241", "", "Eigenmode spatial profile")
	psi_k242		= NewScalarExcitation("psi_k242", "", "Eigenmode spatial profile")
	psi_k243		= NewScalarExcitation("psi_k243", "", "Eigenmode spatial profile")
	psi_k244		= NewScalarExcitation("psi_k244", "", "Eigenmode spatial profile")
	psi_k245		= NewScalarExcitation("psi_k245", "", "Eigenmode spatial profile")
	psi_k246		= NewScalarExcitation("psi_k246", "", "Eigenmode spatial profile")
	psi_k247		= NewScalarExcitation("psi_k247", "", "Eigenmode spatial profile")
	psi_k248		= NewScalarExcitation("psi_k248", "", "Eigenmode spatial profile")
	psi_k249		= NewScalarExcitation("psi_k249", "", "Eigenmode spatial profile")
	psi_k250		= NewScalarExcitation("psi_k250", "", "Eigenmode spatial profile")
	psi_k251		= NewScalarExcitation("psi_k251", "", "Eigenmode spatial profile")
	psi_k252		= NewScalarExcitation("psi_k252", "", "Eigenmode spatial profile")
	psi_k253		= NewScalarExcitation("psi_k253", "", "Eigenmode spatial profile")
	psi_k254		= NewScalarExcitation("psi_k254", "", "Eigenmode spatial profile")
	psi_k255		= NewScalarExcitation("psi_k255", "", "Eigenmode spatial profile")
	psi_k256		= NewScalarExcitation("psi_k256", "", "Eigenmode spatial profile")
	psi_k257		= NewScalarExcitation("psi_k257", "", "Eigenmode spatial profile")
	psi_k258		= NewScalarExcitation("psi_k258", "", "Eigenmode spatial profile")
	psi_k259		= NewScalarExcitation("psi_k259", "", "Eigenmode spatial profile")
	psi_k260		= NewScalarExcitation("psi_k260", "", "Eigenmode spatial profile")
	psi_k261		= NewScalarExcitation("psi_k261", "", "Eigenmode spatial profile")
	psi_k262		= NewScalarExcitation("psi_k262", "", "Eigenmode spatial profile")
	psi_k263		= NewScalarExcitation("psi_k263", "", "Eigenmode spatial profile")
	psi_k264		= NewScalarExcitation("psi_k264", "", "Eigenmode spatial profile")
	psi_k265		= NewScalarExcitation("psi_k265", "", "Eigenmode spatial profile")
	psi_k266		= NewScalarExcitation("psi_k266", "", "Eigenmode spatial profile")
	psi_k267		= NewScalarExcitation("psi_k267", "", "Eigenmode spatial profile")
	psi_k268		= NewScalarExcitation("psi_k268", "", "Eigenmode spatial profile")
	psi_k269		= NewScalarExcitation("psi_k269", "", "Eigenmode spatial profile")
	psi_k270		= NewScalarExcitation("psi_k270", "", "Eigenmode spatial profile")
	psi_k271		= NewScalarExcitation("psi_k271", "", "Eigenmode spatial profile")
	psi_k272		= NewScalarExcitation("psi_k272", "", "Eigenmode spatial profile")
	psi_k273		= NewScalarExcitation("psi_k273", "", "Eigenmode spatial profile")
	psi_k274		= NewScalarExcitation("psi_k274", "", "Eigenmode spatial profile")
	psi_k275		= NewScalarExcitation("psi_k275", "", "Eigenmode spatial profile")
	psi_k276		= NewScalarExcitation("psi_k276", "", "Eigenmode spatial profile")
	psi_k277		= NewScalarExcitation("psi_k277", "", "Eigenmode spatial profile")
	psi_k278		= NewScalarExcitation("psi_k278", "", "Eigenmode spatial profile")
	psi_k279		= NewScalarExcitation("psi_k279", "", "Eigenmode spatial profile")
	psi_k280		= NewScalarExcitation("psi_k280", "", "Eigenmode spatial profile")
	psi_k281		= NewScalarExcitation("psi_k281", "", "Eigenmode spatial profile")
	psi_k282		= NewScalarExcitation("psi_k282", "", "Eigenmode spatial profile")
	psi_k283		= NewScalarExcitation("psi_k283", "", "Eigenmode spatial profile")
	psi_k284		= NewScalarExcitation("psi_k284", "", "Eigenmode spatial profile")
	psi_k285		= NewScalarExcitation("psi_k285", "", "Eigenmode spatial profile")
	psi_k286		= NewScalarExcitation("psi_k286", "", "Eigenmode spatial profile")
	psi_k287		= NewScalarExcitation("psi_k287", "", "Eigenmode spatial profile")
	psi_k288		= NewScalarExcitation("psi_k288", "", "Eigenmode spatial profile")
	psi_k289		= NewScalarExcitation("psi_k289", "", "Eigenmode spatial profile")
	psi_k290		= NewScalarExcitation("psi_k290", "", "Eigenmode spatial profile")
	psi_k291		= NewScalarExcitation("psi_k291", "", "Eigenmode spatial profile")
	psi_k292		= NewScalarExcitation("psi_k292", "", "Eigenmode spatial profile")
	psi_k293		= NewScalarExcitation("psi_k293", "", "Eigenmode spatial profile")
	psi_k294		= NewScalarExcitation("psi_k294", "", "Eigenmode spatial profile")
	psi_k295		= NewScalarExcitation("psi_k295", "", "Eigenmode spatial profile")
	psi_k296		= NewScalarExcitation("psi_k296", "", "Eigenmode spatial profile")
	psi_k297		= NewScalarExcitation("psi_k297", "", "Eigenmode spatial profile")
	psi_k298		= NewScalarExcitation("psi_k298", "", "Eigenmode spatial profile")
	psi_k299		= NewScalarExcitation("psi_k299", "", "Eigenmode spatial profile")
	psi_k300		= NewScalarExcitation("psi_k300", "", "Eigenmode spatial profile")
	psi_k301		= NewScalarExcitation("psi_k301", "", "Eigenmode spatial profile")
	psi_k302		= NewScalarExcitation("psi_k302", "", "Eigenmode spatial profile")
	psi_k303		= NewScalarExcitation("psi_k303", "", "Eigenmode spatial profile")
	psi_k304		= NewScalarExcitation("psi_k304", "", "Eigenmode spatial profile")
	psi_k305		= NewScalarExcitation("psi_k305", "", "Eigenmode spatial profile")
	psi_k306		= NewScalarExcitation("psi_k306", "", "Eigenmode spatial profile")
	psi_k307		= NewScalarExcitation("psi_k307", "", "Eigenmode spatial profile")
	psi_k308		= NewScalarExcitation("psi_k308", "", "Eigenmode spatial profile")
	psi_k309		= NewScalarExcitation("psi_k309", "", "Eigenmode spatial profile")
	psi_k310		= NewScalarExcitation("psi_k310", "", "Eigenmode spatial profile")
	psi_k311		= NewScalarExcitation("psi_k311", "", "Eigenmode spatial profile")
	psi_k312		= NewScalarExcitation("psi_k312", "", "Eigenmode spatial profile")
	psi_k313		= NewScalarExcitation("psi_k313", "", "Eigenmode spatial profile")
	psi_k314		= NewScalarExcitation("psi_k314", "", "Eigenmode spatial profile")
	psi_k315		= NewScalarExcitation("psi_k315", "", "Eigenmode spatial profile")
	psi_k316		= NewScalarExcitation("psi_k316", "", "Eigenmode spatial profile")
	psi_k317		= NewScalarExcitation("psi_k317", "", "Eigenmode spatial profile")
	psi_k318		= NewScalarExcitation("psi_k318", "", "Eigenmode spatial profile")
	psi_k319		= NewScalarExcitation("psi_k319", "", "Eigenmode spatial profile")
	psi_k320		= NewScalarExcitation("psi_k320", "", "Eigenmode spatial profile")
	psi_k321		= NewScalarExcitation("psi_k321", "", "Eigenmode spatial profile")
	psi_k322		= NewScalarExcitation("psi_k322", "", "Eigenmode spatial profile")
	psi_k323		= NewScalarExcitation("psi_k323", "", "Eigenmode spatial profile")
	psi_k324		= NewScalarExcitation("psi_k324", "", "Eigenmode spatial profile")
	psi_k325		= NewScalarExcitation("psi_k325", "", "Eigenmode spatial profile")
	psi_k326		= NewScalarExcitation("psi_k326", "", "Eigenmode spatial profile")
	psi_k327		= NewScalarExcitation("psi_k327", "", "Eigenmode spatial profile")
	psi_k328		= NewScalarExcitation("psi_k328", "", "Eigenmode spatial profile")
	psi_k329		= NewScalarExcitation("psi_k329", "", "Eigenmode spatial profile")
	psi_k330		= NewScalarExcitation("psi_k330", "", "Eigenmode spatial profile")
	psi_k331		= NewScalarExcitation("psi_k331", "", "Eigenmode spatial profile")
	psi_k332		= NewScalarExcitation("psi_k332", "", "Eigenmode spatial profile")
	psi_k333		= NewScalarExcitation("psi_k333", "", "Eigenmode spatial profile")
	psi_k334		= NewScalarExcitation("psi_k334", "", "Eigenmode spatial profile")
	psi_k335		= NewScalarExcitation("psi_k335", "", "Eigenmode spatial profile")
	psi_k336		= NewScalarExcitation("psi_k336", "", "Eigenmode spatial profile")
	psi_k337		= NewScalarExcitation("psi_k337", "", "Eigenmode spatial profile")
	psi_k338		= NewScalarExcitation("psi_k338", "", "Eigenmode spatial profile")
	psi_k339		= NewScalarExcitation("psi_k339", "", "Eigenmode spatial profile")
	psi_k340		= NewScalarExcitation("psi_k340", "", "Eigenmode spatial profile")
	psi_k341		= NewScalarExcitation("psi_k341", "", "Eigenmode spatial profile")
	psi_k342		= NewScalarExcitation("psi_k342", "", "Eigenmode spatial profile")
	psi_k343		= NewScalarExcitation("psi_k343", "", "Eigenmode spatial profile")
	psi_k344		= NewScalarExcitation("psi_k344", "", "Eigenmode spatial profile")
	psi_k345		= NewScalarExcitation("psi_k345", "", "Eigenmode spatial profile")
	psi_k346		= NewScalarExcitation("psi_k346", "", "Eigenmode spatial profile")
	psi_k347		= NewScalarExcitation("psi_k347", "", "Eigenmode spatial profile")
	psi_k348		= NewScalarExcitation("psi_k348", "", "Eigenmode spatial profile")
	psi_k349		= NewScalarExcitation("psi_k349", "", "Eigenmode spatial profile")
	psi_k350		= NewScalarExcitation("psi_k350", "", "Eigenmode spatial profile")
	psi_k351		= NewScalarExcitation("psi_k351", "", "Eigenmode spatial profile")
	psi_k352		= NewScalarExcitation("psi_k352", "", "Eigenmode spatial profile")
	psi_k353		= NewScalarExcitation("psi_k353", "", "Eigenmode spatial profile")
	psi_k354		= NewScalarExcitation("psi_k354", "", "Eigenmode spatial profile")
	psi_k355		= NewScalarExcitation("psi_k355", "", "Eigenmode spatial profile")
	psi_k356		= NewScalarExcitation("psi_k356", "", "Eigenmode spatial profile")
	psi_k357		= NewScalarExcitation("psi_k357", "", "Eigenmode spatial profile")
	psi_k358		= NewScalarExcitation("psi_k358", "", "Eigenmode spatial profile")
	psi_k359		= NewScalarExcitation("psi_k359", "", "Eigenmode spatial profile")
	psi_k360		= NewScalarExcitation("psi_k360", "", "Eigenmode spatial profile")
	psi_k361		= NewScalarExcitation("psi_k361", "", "Eigenmode spatial profile")
	psi_k362		= NewScalarExcitation("psi_k362", "", "Eigenmode spatial profile")
	psi_k363		= NewScalarExcitation("psi_k363", "", "Eigenmode spatial profile")
	psi_k364		= NewScalarExcitation("psi_k364", "", "Eigenmode spatial profile")
	psi_k365		= NewScalarExcitation("psi_k365", "", "Eigenmode spatial profile")
	psi_k366		= NewScalarExcitation("psi_k366", "", "Eigenmode spatial profile")
	psi_k367		= NewScalarExcitation("psi_k367", "", "Eigenmode spatial profile")
	psi_k368		= NewScalarExcitation("psi_k368", "", "Eigenmode spatial profile")
	psi_k369		= NewScalarExcitation("psi_k369", "", "Eigenmode spatial profile")
	psi_k370		= NewScalarExcitation("psi_k370", "", "Eigenmode spatial profile")
	psi_k371		= NewScalarExcitation("psi_k371", "", "Eigenmode spatial profile")
	psi_k372		= NewScalarExcitation("psi_k372", "", "Eigenmode spatial profile")
	psi_k373		= NewScalarExcitation("psi_k373", "", "Eigenmode spatial profile")
	psi_k374		= NewScalarExcitation("psi_k374", "", "Eigenmode spatial profile")
	psi_k375		= NewScalarExcitation("psi_k375", "", "Eigenmode spatial profile")
	psi_k376		= NewScalarExcitation("psi_k376", "", "Eigenmode spatial profile")
	psi_k377		= NewScalarExcitation("psi_k377", "", "Eigenmode spatial profile")
	psi_k378		= NewScalarExcitation("psi_k378", "", "Eigenmode spatial profile")
	psi_k379		= NewScalarExcitation("psi_k379", "", "Eigenmode spatial profile")
	psi_k380		= NewScalarExcitation("psi_k380", "", "Eigenmode spatial profile")
	psi_k381		= NewScalarExcitation("psi_k381", "", "Eigenmode spatial profile")
	psi_k382		= NewScalarExcitation("psi_k382", "", "Eigenmode spatial profile")
	psi_k383		= NewScalarExcitation("psi_k383", "", "Eigenmode spatial profile")
	psi_k384		= NewScalarExcitation("psi_k384", "", "Eigenmode spatial profile")
	psi_k385		= NewScalarExcitation("psi_k385", "", "Eigenmode spatial profile")
	psi_k386		= NewScalarExcitation("psi_k386", "", "Eigenmode spatial profile")
	psi_k387		= NewScalarExcitation("psi_k387", "", "Eigenmode spatial profile")
	psi_k388		= NewScalarExcitation("psi_k388", "", "Eigenmode spatial profile")
	psi_k389		= NewScalarExcitation("psi_k389", "", "Eigenmode spatial profile")
	psi_k390		= NewScalarExcitation("psi_k390", "", "Eigenmode spatial profile")
	psi_k391		= NewScalarExcitation("psi_k391", "", "Eigenmode spatial profile")
	psi_k392		= NewScalarExcitation("psi_k392", "", "Eigenmode spatial profile")
	psi_k393		= NewScalarExcitation("psi_k393", "", "Eigenmode spatial profile")
	psi_k394		= NewScalarExcitation("psi_k394", "", "Eigenmode spatial profile")
	psi_k395		= NewScalarExcitation("psi_k395", "", "Eigenmode spatial profile")
	psi_k396		= NewScalarExcitation("psi_k396", "", "Eigenmode spatial profile")
	psi_k397		= NewScalarExcitation("psi_k397", "", "Eigenmode spatial profile")
	psi_k398		= NewScalarExcitation("psi_k398", "", "Eigenmode spatial profile")
	psi_k399		= NewScalarExcitation("psi_k399", "", "Eigenmode spatial profile")
	psi_k400		= NewScalarExcitation("psi_k400", "", "Eigenmode spatial profile")
	psi_k401		= NewScalarExcitation("psi_k401", "", "Eigenmode spatial profile")
	psi_k402		= NewScalarExcitation("psi_k402", "", "Eigenmode spatial profile")
	psi_k403		= NewScalarExcitation("psi_k403", "", "Eigenmode spatial profile")
	psi_k404		= NewScalarExcitation("psi_k404", "", "Eigenmode spatial profile")
	psi_k405		= NewScalarExcitation("psi_k405", "", "Eigenmode spatial profile")
	psi_k406		= NewScalarExcitation("psi_k406", "", "Eigenmode spatial profile")
	psi_k407		= NewScalarExcitation("psi_k407", "", "Eigenmode spatial profile")
	psi_k408		= NewScalarExcitation("psi_k408", "", "Eigenmode spatial profile")
	psi_k409		= NewScalarExcitation("psi_k409", "", "Eigenmode spatial profile")
	psi_k410		= NewScalarExcitation("psi_k410", "", "Eigenmode spatial profile")
	psi_k411		= NewScalarExcitation("psi_k411", "", "Eigenmode spatial profile")
	psi_k412		= NewScalarExcitation("psi_k412", "", "Eigenmode spatial profile")
	psi_k413		= NewScalarExcitation("psi_k413", "", "Eigenmode spatial profile")
	psi_k414		= NewScalarExcitation("psi_k414", "", "Eigenmode spatial profile")
	psi_k415		= NewScalarExcitation("psi_k415", "", "Eigenmode spatial profile")
	psi_k416		= NewScalarExcitation("psi_k416", "", "Eigenmode spatial profile")
	psi_k417		= NewScalarExcitation("psi_k417", "", "Eigenmode spatial profile")
	psi_k418		= NewScalarExcitation("psi_k418", "", "Eigenmode spatial profile")
	psi_k419		= NewScalarExcitation("psi_k419", "", "Eigenmode spatial profile")
	psi_k420		= NewScalarExcitation("psi_k420", "", "Eigenmode spatial profile")
	psi_k421		= NewScalarExcitation("psi_k421", "", "Eigenmode spatial profile")
	psi_k422		= NewScalarExcitation("psi_k422", "", "Eigenmode spatial profile")
	psi_k423		= NewScalarExcitation("psi_k423", "", "Eigenmode spatial profile")
	psi_k424		= NewScalarExcitation("psi_k424", "", "Eigenmode spatial profile")
	psi_k425		= NewScalarExcitation("psi_k425", "", "Eigenmode spatial profile")
	psi_k426		= NewScalarExcitation("psi_k426", "", "Eigenmode spatial profile")
	psi_k427		= NewScalarExcitation("psi_k427", "", "Eigenmode spatial profile")
	psi_k428		= NewScalarExcitation("psi_k428", "", "Eigenmode spatial profile")
	psi_k429		= NewScalarExcitation("psi_k429", "", "Eigenmode spatial profile")
	psi_k430		= NewScalarExcitation("psi_k430", "", "Eigenmode spatial profile")
	psi_k431		= NewScalarExcitation("psi_k431", "", "Eigenmode spatial profile")
	psi_k432		= NewScalarExcitation("psi_k432", "", "Eigenmode spatial profile")
	psi_k433		= NewScalarExcitation("psi_k433", "", "Eigenmode spatial profile")
	psi_k434		= NewScalarExcitation("psi_k434", "", "Eigenmode spatial profile")
	psi_k435		= NewScalarExcitation("psi_k435", "", "Eigenmode spatial profile")
	psi_k436		= NewScalarExcitation("psi_k436", "", "Eigenmode spatial profile")
	psi_k437		= NewScalarExcitation("psi_k437", "", "Eigenmode spatial profile")
	psi_k438		= NewScalarExcitation("psi_k438", "", "Eigenmode spatial profile")
	psi_k439		= NewScalarExcitation("psi_k439", "", "Eigenmode spatial profile")
	psi_k440		= NewScalarExcitation("psi_k440", "", "Eigenmode spatial profile")
	psi_k441		= NewScalarExcitation("psi_k441", "", "Eigenmode spatial profile")
	psi_k442		= NewScalarExcitation("psi_k442", "", "Eigenmode spatial profile")
	psi_k443		= NewScalarExcitation("psi_k443", "", "Eigenmode spatial profile")
	psi_k444		= NewScalarExcitation("psi_k444", "", "Eigenmode spatial profile")
	psi_k445		= NewScalarExcitation("psi_k445", "", "Eigenmode spatial profile")
	psi_k446		= NewScalarExcitation("psi_k446", "", "Eigenmode spatial profile")
	psi_k447		= NewScalarExcitation("psi_k447", "", "Eigenmode spatial profile")
	psi_k448		= NewScalarExcitation("psi_k448", "", "Eigenmode spatial profile")
	psi_k449		= NewScalarExcitation("psi_k449", "", "Eigenmode spatial profile")
	psi_k450		= NewScalarExcitation("psi_k450", "", "Eigenmode spatial profile")
	psi_k451		= NewScalarExcitation("psi_k451", "", "Eigenmode spatial profile")
	psi_k452		= NewScalarExcitation("psi_k452", "", "Eigenmode spatial profile")
	psi_k453		= NewScalarExcitation("psi_k453", "", "Eigenmode spatial profile")
	psi_k454		= NewScalarExcitation("psi_k454", "", "Eigenmode spatial profile")
	psi_k455		= NewScalarExcitation("psi_k455", "", "Eigenmode spatial profile")
	psi_k456		= NewScalarExcitation("psi_k456", "", "Eigenmode spatial profile")
	psi_k457		= NewScalarExcitation("psi_k457", "", "Eigenmode spatial profile")
	psi_k458		= NewScalarExcitation("psi_k458", "", "Eigenmode spatial profile")
	psi_k459		= NewScalarExcitation("psi_k459", "", "Eigenmode spatial profile")
	psi_k460		= NewScalarExcitation("psi_k460", "", "Eigenmode spatial profile")
	psi_k461		= NewScalarExcitation("psi_k461", "", "Eigenmode spatial profile")
	psi_k462		= NewScalarExcitation("psi_k462", "", "Eigenmode spatial profile")
	psi_k463		= NewScalarExcitation("psi_k463", "", "Eigenmode spatial profile")
	psi_k464		= NewScalarExcitation("psi_k464", "", "Eigenmode spatial profile")
	psi_k465		= NewScalarExcitation("psi_k465", "", "Eigenmode spatial profile")
	psi_k466		= NewScalarExcitation("psi_k466", "", "Eigenmode spatial profile")
	psi_k467		= NewScalarExcitation("psi_k467", "", "Eigenmode spatial profile")
	psi_k468		= NewScalarExcitation("psi_k468", "", "Eigenmode spatial profile")
	psi_k469		= NewScalarExcitation("psi_k469", "", "Eigenmode spatial profile")
	psi_k470		= NewScalarExcitation("psi_k470", "", "Eigenmode spatial profile")
	psi_k471		= NewScalarExcitation("psi_k471", "", "Eigenmode spatial profile")
	psi_k472		= NewScalarExcitation("psi_k472", "", "Eigenmode spatial profile")
	psi_k473		= NewScalarExcitation("psi_k473", "", "Eigenmode spatial profile")
	psi_k474		= NewScalarExcitation("psi_k474", "", "Eigenmode spatial profile")
	psi_k475		= NewScalarExcitation("psi_k475", "", "Eigenmode spatial profile")
	psi_k476		= NewScalarExcitation("psi_k476", "", "Eigenmode spatial profile")
	psi_k477		= NewScalarExcitation("psi_k477", "", "Eigenmode spatial profile")
	psi_k478		= NewScalarExcitation("psi_k478", "", "Eigenmode spatial profile")
	psi_k479		= NewScalarExcitation("psi_k479", "", "Eigenmode spatial profile")
	psi_k480		= NewScalarExcitation("psi_k480", "", "Eigenmode spatial profile")
	psi_k481		= NewScalarExcitation("psi_k481", "", "Eigenmode spatial profile")
	psi_k482		= NewScalarExcitation("psi_k482", "", "Eigenmode spatial profile")
	psi_k483		= NewScalarExcitation("psi_k483", "", "Eigenmode spatial profile")
	psi_k484		= NewScalarExcitation("psi_k484", "", "Eigenmode spatial profile")
	psi_k485		= NewScalarExcitation("psi_k485", "", "Eigenmode spatial profile")
	psi_k486		= NewScalarExcitation("psi_k486", "", "Eigenmode spatial profile")
	psi_k487		= NewScalarExcitation("psi_k487", "", "Eigenmode spatial profile")
	psi_k488		= NewScalarExcitation("psi_k488", "", "Eigenmode spatial profile")
	psi_k489		= NewScalarExcitation("psi_k489", "", "Eigenmode spatial profile")
	psi_k490		= NewScalarExcitation("psi_k490", "", "Eigenmode spatial profile")
	psi_k491		= NewScalarExcitation("psi_k491", "", "Eigenmode spatial profile")
	psi_k492		= NewScalarExcitation("psi_k492", "", "Eigenmode spatial profile")
	psi_k493		= NewScalarExcitation("psi_k493", "", "Eigenmode spatial profile")
	psi_k494		= NewScalarExcitation("psi_k494", "", "Eigenmode spatial profile")
	psi_k495		= NewScalarExcitation("psi_k495", "", "Eigenmode spatial profile")
	psi_k496		= NewScalarExcitation("psi_k496", "", "Eigenmode spatial profile")
	psi_k497		= NewScalarExcitation("psi_k497", "", "Eigenmode spatial profile")
	psi_k498		= NewScalarExcitation("psi_k498", "", "Eigenmode spatial profile")
	psi_k499		= NewScalarExcitation("psi_k499", "", "Eigenmode spatial profile")
	
			
	a_k000		= NewVectorValue("a_k000", "", "delta_mxy projection onto psi_k000", GetModeAmplitudek000 )
	a_k001		= NewVectorValue("a_k001", "", "delta_mxy projection onto psi_k001", GetModeAmplitudek001 )
	a_k002		= NewVectorValue("a_k002", "", "delta_mxy projection onto psi_k002", GetModeAmplitudek002 )
	a_k003		= NewVectorValue("a_k003", "", "delta_mxy projection onto psi_k003", GetModeAmplitudek003 )
	a_k004		= NewVectorValue("a_k004", "", "delta_mxy projection onto psi_k004", GetModeAmplitudek004 )
	a_k005		= NewVectorValue("a_k005", "", "delta_mxy projection onto psi_k005", GetModeAmplitudek005 )
	a_k006		= NewVectorValue("a_k006", "", "delta_mxy projection onto psi_k006", GetModeAmplitudek006 )
	a_k007		= NewVectorValue("a_k007", "", "delta_mxy projection onto psi_k007", GetModeAmplitudek007 )
	a_k008		= NewVectorValue("a_k008", "", "delta_mxy projection onto psi_k008", GetModeAmplitudek008 )
	a_k009		= NewVectorValue("a_k009", "", "delta_mxy projection onto psi_k009", GetModeAmplitudek009 )
	a_k010		= NewVectorValue("a_k010", "", "delta_mxy projection onto psi_k010", GetModeAmplitudek010 )
	a_k011		= NewVectorValue("a_k011", "", "delta_mxy projection onto psi_k011", GetModeAmplitudek011 )
	a_k012		= NewVectorValue("a_k012", "", "delta_mxy projection onto psi_k012", GetModeAmplitudek012 )
	a_k013		= NewVectorValue("a_k013", "", "delta_mxy projection onto psi_k013", GetModeAmplitudek013 )
	a_k014		= NewVectorValue("a_k014", "", "delta_mxy projection onto psi_k014", GetModeAmplitudek014 )
	a_k015		= NewVectorValue("a_k015", "", "delta_mxy projection onto psi_k015", GetModeAmplitudek015 )
	a_k016		= NewVectorValue("a_k016", "", "delta_mxy projection onto psi_k016", GetModeAmplitudek016 )
	a_k017		= NewVectorValue("a_k017", "", "delta_mxy projection onto psi_k017", GetModeAmplitudek017 )
	a_k018		= NewVectorValue("a_k018", "", "delta_mxy projection onto psi_k018", GetModeAmplitudek018 )
	a_k019		= NewVectorValue("a_k019", "", "delta_mxy projection onto psi_k019", GetModeAmplitudek019 )
	a_k020		= NewVectorValue("a_k020", "", "delta_mxy projection onto psi_k020", GetModeAmplitudek020 )
	a_k021		= NewVectorValue("a_k021", "", "delta_mxy projection onto psi_k021", GetModeAmplitudek021 )
	a_k022		= NewVectorValue("a_k022", "", "delta_mxy projection onto psi_k022", GetModeAmplitudek022 )
	a_k023		= NewVectorValue("a_k023", "", "delta_mxy projection onto psi_k023", GetModeAmplitudek023 )
	a_k024		= NewVectorValue("a_k024", "", "delta_mxy projection onto psi_k024", GetModeAmplitudek024 )
	a_k025		= NewVectorValue("a_k025", "", "delta_mxy projection onto psi_k025", GetModeAmplitudek025 )
	a_k026		= NewVectorValue("a_k026", "", "delta_mxy projection onto psi_k026", GetModeAmplitudek026 )
	a_k027		= NewVectorValue("a_k027", "", "delta_mxy projection onto psi_k027", GetModeAmplitudek027 )
	a_k028		= NewVectorValue("a_k028", "", "delta_mxy projection onto psi_k028", GetModeAmplitudek028 )
	a_k029		= NewVectorValue("a_k029", "", "delta_mxy projection onto psi_k029", GetModeAmplitudek029 )
	a_k030		= NewVectorValue("a_k030", "", "delta_mxy projection onto psi_k030", GetModeAmplitudek030 )
	a_k031		= NewVectorValue("a_k031", "", "delta_mxy projection onto psi_k031", GetModeAmplitudek031 )
	a_k032		= NewVectorValue("a_k032", "", "delta_mxy projection onto psi_k032", GetModeAmplitudek032 )
	a_k033		= NewVectorValue("a_k033", "", "delta_mxy projection onto psi_k033", GetModeAmplitudek033 )
	a_k034		= NewVectorValue("a_k034", "", "delta_mxy projection onto psi_k034", GetModeAmplitudek034 )
	a_k035		= NewVectorValue("a_k035", "", "delta_mxy projection onto psi_k035", GetModeAmplitudek035 )
	a_k036		= NewVectorValue("a_k036", "", "delta_mxy projection onto psi_k036", GetModeAmplitudek036 )
	a_k037		= NewVectorValue("a_k037", "", "delta_mxy projection onto psi_k037", GetModeAmplitudek037 )
	a_k038		= NewVectorValue("a_k038", "", "delta_mxy projection onto psi_k038", GetModeAmplitudek038 )
	a_k039		= NewVectorValue("a_k039", "", "delta_mxy projection onto psi_k039", GetModeAmplitudek039 )
	a_k040		= NewVectorValue("a_k040", "", "delta_mxy projection onto psi_k040", GetModeAmplitudek040 )
	a_k041		= NewVectorValue("a_k041", "", "delta_mxy projection onto psi_k041", GetModeAmplitudek041 )
	a_k042		= NewVectorValue("a_k042", "", "delta_mxy projection onto psi_k042", GetModeAmplitudek042 )
	a_k043		= NewVectorValue("a_k043", "", "delta_mxy projection onto psi_k043", GetModeAmplitudek043 )
	a_k044		= NewVectorValue("a_k044", "", "delta_mxy projection onto psi_k044", GetModeAmplitudek044 )
	a_k045		= NewVectorValue("a_k045", "", "delta_mxy projection onto psi_k045", GetModeAmplitudek045 )
	a_k046		= NewVectorValue("a_k046", "", "delta_mxy projection onto psi_k046", GetModeAmplitudek046 )
	a_k047		= NewVectorValue("a_k047", "", "delta_mxy projection onto psi_k047", GetModeAmplitudek047 )
	a_k048		= NewVectorValue("a_k048", "", "delta_mxy projection onto psi_k048", GetModeAmplitudek048 )
	a_k049		= NewVectorValue("a_k049", "", "delta_mxy projection onto psi_k049", GetModeAmplitudek049 )
	a_k050		= NewVectorValue("a_k050", "", "delta_mxy projection onto psi_k050", GetModeAmplitudek050 )
	a_k051		= NewVectorValue("a_k051", "", "delta_mxy projection onto psi_k051", GetModeAmplitudek051 )
	a_k052		= NewVectorValue("a_k052", "", "delta_mxy projection onto psi_k052", GetModeAmplitudek052 )
	a_k053		= NewVectorValue("a_k053", "", "delta_mxy projection onto psi_k053", GetModeAmplitudek053 )
	a_k054		= NewVectorValue("a_k054", "", "delta_mxy projection onto psi_k054", GetModeAmplitudek054 )
	a_k055		= NewVectorValue("a_k055", "", "delta_mxy projection onto psi_k055", GetModeAmplitudek055 )
	a_k056		= NewVectorValue("a_k056", "", "delta_mxy projection onto psi_k056", GetModeAmplitudek056 )
	a_k057		= NewVectorValue("a_k057", "", "delta_mxy projection onto psi_k057", GetModeAmplitudek057 )
	a_k058		= NewVectorValue("a_k058", "", "delta_mxy projection onto psi_k058", GetModeAmplitudek058 )
	a_k059		= NewVectorValue("a_k059", "", "delta_mxy projection onto psi_k059", GetModeAmplitudek059 )
	a_k060		= NewVectorValue("a_k060", "", "delta_mxy projection onto psi_k060", GetModeAmplitudek060 )
	a_k061		= NewVectorValue("a_k061", "", "delta_mxy projection onto psi_k061", GetModeAmplitudek061 )
	a_k062		= NewVectorValue("a_k062", "", "delta_mxy projection onto psi_k062", GetModeAmplitudek062 )
	a_k063		= NewVectorValue("a_k063", "", "delta_mxy projection onto psi_k063", GetModeAmplitudek063 )
	a_k064		= NewVectorValue("a_k064", "", "delta_mxy projection onto psi_k064", GetModeAmplitudek064 )
	a_k065		= NewVectorValue("a_k065", "", "delta_mxy projection onto psi_k065", GetModeAmplitudek065 )
	a_k066		= NewVectorValue("a_k066", "", "delta_mxy projection onto psi_k066", GetModeAmplitudek066 )
	a_k067		= NewVectorValue("a_k067", "", "delta_mxy projection onto psi_k067", GetModeAmplitudek067 )
	a_k068		= NewVectorValue("a_k068", "", "delta_mxy projection onto psi_k068", GetModeAmplitudek068 )
	a_k069		= NewVectorValue("a_k069", "", "delta_mxy projection onto psi_k069", GetModeAmplitudek069 )
	a_k070		= NewVectorValue("a_k070", "", "delta_mxy projection onto psi_k070", GetModeAmplitudek070 )
	a_k071		= NewVectorValue("a_k071", "", "delta_mxy projection onto psi_k071", GetModeAmplitudek071 )
	a_k072		= NewVectorValue("a_k072", "", "delta_mxy projection onto psi_k072", GetModeAmplitudek072 )
	a_k073		= NewVectorValue("a_k073", "", "delta_mxy projection onto psi_k073", GetModeAmplitudek073 )
	a_k074		= NewVectorValue("a_k074", "", "delta_mxy projection onto psi_k074", GetModeAmplitudek074 )
	a_k075		= NewVectorValue("a_k075", "", "delta_mxy projection onto psi_k075", GetModeAmplitudek075 )
	a_k076		= NewVectorValue("a_k076", "", "delta_mxy projection onto psi_k076", GetModeAmplitudek076 )
	a_k077		= NewVectorValue("a_k077", "", "delta_mxy projection onto psi_k077", GetModeAmplitudek077 )
	a_k078		= NewVectorValue("a_k078", "", "delta_mxy projection onto psi_k078", GetModeAmplitudek078 )
	a_k079		= NewVectorValue("a_k079", "", "delta_mxy projection onto psi_k079", GetModeAmplitudek079 )
	a_k080		= NewVectorValue("a_k080", "", "delta_mxy projection onto psi_k080", GetModeAmplitudek080 )
	a_k081		= NewVectorValue("a_k081", "", "delta_mxy projection onto psi_k081", GetModeAmplitudek081 )
	a_k082		= NewVectorValue("a_k082", "", "delta_mxy projection onto psi_k082", GetModeAmplitudek082 )
	a_k083		= NewVectorValue("a_k083", "", "delta_mxy projection onto psi_k083", GetModeAmplitudek083 )
	a_k084		= NewVectorValue("a_k084", "", "delta_mxy projection onto psi_k084", GetModeAmplitudek084 )
	a_k085		= NewVectorValue("a_k085", "", "delta_mxy projection onto psi_k085", GetModeAmplitudek085 )
	a_k086		= NewVectorValue("a_k086", "", "delta_mxy projection onto psi_k086", GetModeAmplitudek086 )
	a_k087		= NewVectorValue("a_k087", "", "delta_mxy projection onto psi_k087", GetModeAmplitudek087 )
	a_k088		= NewVectorValue("a_k088", "", "delta_mxy projection onto psi_k088", GetModeAmplitudek088 )
	a_k089		= NewVectorValue("a_k089", "", "delta_mxy projection onto psi_k089", GetModeAmplitudek089 )
	a_k090		= NewVectorValue("a_k090", "", "delta_mxy projection onto psi_k090", GetModeAmplitudek090 )
	a_k091		= NewVectorValue("a_k091", "", "delta_mxy projection onto psi_k091", GetModeAmplitudek091 )
	a_k092		= NewVectorValue("a_k092", "", "delta_mxy projection onto psi_k092", GetModeAmplitudek092 )
	a_k093		= NewVectorValue("a_k093", "", "delta_mxy projection onto psi_k093", GetModeAmplitudek093 )
	a_k094		= NewVectorValue("a_k094", "", "delta_mxy projection onto psi_k094", GetModeAmplitudek094 )
	a_k095		= NewVectorValue("a_k095", "", "delta_mxy projection onto psi_k095", GetModeAmplitudek095 )
	a_k096		= NewVectorValue("a_k096", "", "delta_mxy projection onto psi_k096", GetModeAmplitudek096 )
	a_k097		= NewVectorValue("a_k097", "", "delta_mxy projection onto psi_k097", GetModeAmplitudek097 )
	a_k098		= NewVectorValue("a_k098", "", "delta_mxy projection onto psi_k098", GetModeAmplitudek098 )
	a_k099		= NewVectorValue("a_k099", "", "delta_mxy projection onto psi_k099", GetModeAmplitudek099 )
	a_k100		= NewVectorValue("a_k100", "", "delta_mxy projection onto psi_k100", GetModeAmplitudek100 )
	a_k101		= NewVectorValue("a_k101", "", "delta_mxy projection onto psi_k101", GetModeAmplitudek101 )
	a_k102		= NewVectorValue("a_k102", "", "delta_mxy projection onto psi_k102", GetModeAmplitudek102 )
	a_k103		= NewVectorValue("a_k103", "", "delta_mxy projection onto psi_k103", GetModeAmplitudek103 )
	a_k104		= NewVectorValue("a_k104", "", "delta_mxy projection onto psi_k104", GetModeAmplitudek104 )
	a_k105		= NewVectorValue("a_k105", "", "delta_mxy projection onto psi_k105", GetModeAmplitudek105 )
	a_k106		= NewVectorValue("a_k106", "", "delta_mxy projection onto psi_k106", GetModeAmplitudek106 )
	a_k107		= NewVectorValue("a_k107", "", "delta_mxy projection onto psi_k107", GetModeAmplitudek107 )
	a_k108		= NewVectorValue("a_k108", "", "delta_mxy projection onto psi_k108", GetModeAmplitudek108 )
	a_k109		= NewVectorValue("a_k109", "", "delta_mxy projection onto psi_k109", GetModeAmplitudek109 )
	a_k110		= NewVectorValue("a_k110", "", "delta_mxy projection onto psi_k110", GetModeAmplitudek110 )
	a_k111		= NewVectorValue("a_k111", "", "delta_mxy projection onto psi_k111", GetModeAmplitudek111 )
	a_k112		= NewVectorValue("a_k112", "", "delta_mxy projection onto psi_k112", GetModeAmplitudek112 )
	a_k113		= NewVectorValue("a_k113", "", "delta_mxy projection onto psi_k113", GetModeAmplitudek113 )
	a_k114		= NewVectorValue("a_k114", "", "delta_mxy projection onto psi_k114", GetModeAmplitudek114 )
	a_k115		= NewVectorValue("a_k115", "", "delta_mxy projection onto psi_k115", GetModeAmplitudek115 )
	a_k116		= NewVectorValue("a_k116", "", "delta_mxy projection onto psi_k116", GetModeAmplitudek116 )
	a_k117		= NewVectorValue("a_k117", "", "delta_mxy projection onto psi_k117", GetModeAmplitudek117 )
	a_k118		= NewVectorValue("a_k118", "", "delta_mxy projection onto psi_k118", GetModeAmplitudek118 )
	a_k119		= NewVectorValue("a_k119", "", "delta_mxy projection onto psi_k119", GetModeAmplitudek119 )
	a_k120		= NewVectorValue("a_k120", "", "delta_mxy projection onto psi_k120", GetModeAmplitudek120 )
	a_k121		= NewVectorValue("a_k121", "", "delta_mxy projection onto psi_k121", GetModeAmplitudek121 )
	a_k122		= NewVectorValue("a_k122", "", "delta_mxy projection onto psi_k122", GetModeAmplitudek122 )
	a_k123		= NewVectorValue("a_k123", "", "delta_mxy projection onto psi_k123", GetModeAmplitudek123 )
	a_k124		= NewVectorValue("a_k124", "", "delta_mxy projection onto psi_k124", GetModeAmplitudek124 )
	a_k125		= NewVectorValue("a_k125", "", "delta_mxy projection onto psi_k125", GetModeAmplitudek125 )
	a_k126		= NewVectorValue("a_k126", "", "delta_mxy projection onto psi_k126", GetModeAmplitudek126 )
	a_k127		= NewVectorValue("a_k127", "", "delta_mxy projection onto psi_k127", GetModeAmplitudek127 )
	a_k128		= NewVectorValue("a_k128", "", "delta_mxy projection onto psi_k128", GetModeAmplitudek128 )
	a_k129		= NewVectorValue("a_k129", "", "delta_mxy projection onto psi_k129", GetModeAmplitudek129 )
	a_k130		= NewVectorValue("a_k130", "", "delta_mxy projection onto psi_k130", GetModeAmplitudek130 )
	a_k131		= NewVectorValue("a_k131", "", "delta_mxy projection onto psi_k131", GetModeAmplitudek131 )
	a_k132		= NewVectorValue("a_k132", "", "delta_mxy projection onto psi_k132", GetModeAmplitudek132 )
	a_k133		= NewVectorValue("a_k133", "", "delta_mxy projection onto psi_k133", GetModeAmplitudek133 )
	a_k134		= NewVectorValue("a_k134", "", "delta_mxy projection onto psi_k134", GetModeAmplitudek134 )
	a_k135		= NewVectorValue("a_k135", "", "delta_mxy projection onto psi_k135", GetModeAmplitudek135 )
	a_k136		= NewVectorValue("a_k136", "", "delta_mxy projection onto psi_k136", GetModeAmplitudek136 )
	a_k137		= NewVectorValue("a_k137", "", "delta_mxy projection onto psi_k137", GetModeAmplitudek137 )
	a_k138		= NewVectorValue("a_k138", "", "delta_mxy projection onto psi_k138", GetModeAmplitudek138 )
	a_k139		= NewVectorValue("a_k139", "", "delta_mxy projection onto psi_k139", GetModeAmplitudek139 )
	a_k140		= NewVectorValue("a_k140", "", "delta_mxy projection onto psi_k140", GetModeAmplitudek140 )
	a_k141		= NewVectorValue("a_k141", "", "delta_mxy projection onto psi_k141", GetModeAmplitudek141 )
	a_k142		= NewVectorValue("a_k142", "", "delta_mxy projection onto psi_k142", GetModeAmplitudek142 )
	a_k143		= NewVectorValue("a_k143", "", "delta_mxy projection onto psi_k143", GetModeAmplitudek143 )
	a_k144		= NewVectorValue("a_k144", "", "delta_mxy projection onto psi_k144", GetModeAmplitudek144 )
	a_k145		= NewVectorValue("a_k145", "", "delta_mxy projection onto psi_k145", GetModeAmplitudek145 )
	a_k146		= NewVectorValue("a_k146", "", "delta_mxy projection onto psi_k146", GetModeAmplitudek146 )
	a_k147		= NewVectorValue("a_k147", "", "delta_mxy projection onto psi_k147", GetModeAmplitudek147 )
	a_k148		= NewVectorValue("a_k148", "", "delta_mxy projection onto psi_k148", GetModeAmplitudek148 )
	a_k149		= NewVectorValue("a_k149", "", "delta_mxy projection onto psi_k149", GetModeAmplitudek149 )
	a_k150		= NewVectorValue("a_k150", "", "delta_mxy projection onto psi_k150", GetModeAmplitudek150 )
	a_k151		= NewVectorValue("a_k151", "", "delta_mxy projection onto psi_k151", GetModeAmplitudek151 )
	a_k152		= NewVectorValue("a_k152", "", "delta_mxy projection onto psi_k152", GetModeAmplitudek152 )
	a_k153		= NewVectorValue("a_k153", "", "delta_mxy projection onto psi_k153", GetModeAmplitudek153 )
	a_k154		= NewVectorValue("a_k154", "", "delta_mxy projection onto psi_k154", GetModeAmplitudek154 )
	a_k155		= NewVectorValue("a_k155", "", "delta_mxy projection onto psi_k155", GetModeAmplitudek155 )
	a_k156		= NewVectorValue("a_k156", "", "delta_mxy projection onto psi_k156", GetModeAmplitudek156 )
	a_k157		= NewVectorValue("a_k157", "", "delta_mxy projection onto psi_k157", GetModeAmplitudek157 )
	a_k158		= NewVectorValue("a_k158", "", "delta_mxy projection onto psi_k158", GetModeAmplitudek158 )
	a_k159		= NewVectorValue("a_k159", "", "delta_mxy projection onto psi_k159", GetModeAmplitudek159 )
	a_k160		= NewVectorValue("a_k160", "", "delta_mxy projection onto psi_k160", GetModeAmplitudek160 )
	a_k161		= NewVectorValue("a_k161", "", "delta_mxy projection onto psi_k161", GetModeAmplitudek161 )
	a_k162		= NewVectorValue("a_k162", "", "delta_mxy projection onto psi_k162", GetModeAmplitudek162 )
	a_k163		= NewVectorValue("a_k163", "", "delta_mxy projection onto psi_k163", GetModeAmplitudek163 )
	a_k164		= NewVectorValue("a_k164", "", "delta_mxy projection onto psi_k164", GetModeAmplitudek164 )
	a_k165		= NewVectorValue("a_k165", "", "delta_mxy projection onto psi_k165", GetModeAmplitudek165 )
	a_k166		= NewVectorValue("a_k166", "", "delta_mxy projection onto psi_k166", GetModeAmplitudek166 )
	a_k167		= NewVectorValue("a_k167", "", "delta_mxy projection onto psi_k167", GetModeAmplitudek167 )
	a_k168		= NewVectorValue("a_k168", "", "delta_mxy projection onto psi_k168", GetModeAmplitudek168 )
	a_k169		= NewVectorValue("a_k169", "", "delta_mxy projection onto psi_k169", GetModeAmplitudek169 )
	a_k170		= NewVectorValue("a_k170", "", "delta_mxy projection onto psi_k170", GetModeAmplitudek170 )
	a_k171		= NewVectorValue("a_k171", "", "delta_mxy projection onto psi_k171", GetModeAmplitudek171 )
	a_k172		= NewVectorValue("a_k172", "", "delta_mxy projection onto psi_k172", GetModeAmplitudek172 )
	a_k173		= NewVectorValue("a_k173", "", "delta_mxy projection onto psi_k173", GetModeAmplitudek173 )
	a_k174		= NewVectorValue("a_k174", "", "delta_mxy projection onto psi_k174", GetModeAmplitudek174 )
	a_k175		= NewVectorValue("a_k175", "", "delta_mxy projection onto psi_k175", GetModeAmplitudek175 )
	a_k176		= NewVectorValue("a_k176", "", "delta_mxy projection onto psi_k176", GetModeAmplitudek176 )
	a_k177		= NewVectorValue("a_k177", "", "delta_mxy projection onto psi_k177", GetModeAmplitudek177 )
	a_k178		= NewVectorValue("a_k178", "", "delta_mxy projection onto psi_k178", GetModeAmplitudek178 )
	a_k179		= NewVectorValue("a_k179", "", "delta_mxy projection onto psi_k179", GetModeAmplitudek179 )
	a_k180		= NewVectorValue("a_k180", "", "delta_mxy projection onto psi_k180", GetModeAmplitudek180 )
	a_k181		= NewVectorValue("a_k181", "", "delta_mxy projection onto psi_k181", GetModeAmplitudek181 )
	a_k182		= NewVectorValue("a_k182", "", "delta_mxy projection onto psi_k182", GetModeAmplitudek182 )
	a_k183		= NewVectorValue("a_k183", "", "delta_mxy projection onto psi_k183", GetModeAmplitudek183 )
	a_k184		= NewVectorValue("a_k184", "", "delta_mxy projection onto psi_k184", GetModeAmplitudek184 )
	a_k185		= NewVectorValue("a_k185", "", "delta_mxy projection onto psi_k185", GetModeAmplitudek185 )
	a_k186		= NewVectorValue("a_k186", "", "delta_mxy projection onto psi_k186", GetModeAmplitudek186 )
	a_k187		= NewVectorValue("a_k187", "", "delta_mxy projection onto psi_k187", GetModeAmplitudek187 )
	a_k188		= NewVectorValue("a_k188", "", "delta_mxy projection onto psi_k188", GetModeAmplitudek188 )
	a_k189		= NewVectorValue("a_k189", "", "delta_mxy projection onto psi_k189", GetModeAmplitudek189 )
	a_k190		= NewVectorValue("a_k190", "", "delta_mxy projection onto psi_k190", GetModeAmplitudek190 )
	a_k191		= NewVectorValue("a_k191", "", "delta_mxy projection onto psi_k191", GetModeAmplitudek191 )
	a_k192		= NewVectorValue("a_k192", "", "delta_mxy projection onto psi_k192", GetModeAmplitudek192 )
	a_k193		= NewVectorValue("a_k193", "", "delta_mxy projection onto psi_k193", GetModeAmplitudek193 )
	a_k194		= NewVectorValue("a_k194", "", "delta_mxy projection onto psi_k194", GetModeAmplitudek194 )
	a_k195		= NewVectorValue("a_k195", "", "delta_mxy projection onto psi_k195", GetModeAmplitudek195 )
	a_k196		= NewVectorValue("a_k196", "", "delta_mxy projection onto psi_k196", GetModeAmplitudek196 )
	a_k197		= NewVectorValue("a_k197", "", "delta_mxy projection onto psi_k197", GetModeAmplitudek197 )
	a_k198		= NewVectorValue("a_k198", "", "delta_mxy projection onto psi_k198", GetModeAmplitudek198 )
	a_k199		= NewVectorValue("a_k199", "", "delta_mxy projection onto psi_k199", GetModeAmplitudek199 )
	a_k200		= NewVectorValue("a_k200", "", "delta_mxy projection onto psi_k200", GetModeAmplitudek200 )
	a_k201		= NewVectorValue("a_k201", "", "delta_mxy projection onto psi_k201", GetModeAmplitudek201 )
	a_k202		= NewVectorValue("a_k202", "", "delta_mxy projection onto psi_k202", GetModeAmplitudek202 )
	a_k203		= NewVectorValue("a_k203", "", "delta_mxy projection onto psi_k203", GetModeAmplitudek203 )
	a_k204		= NewVectorValue("a_k204", "", "delta_mxy projection onto psi_k204", GetModeAmplitudek204 )
	a_k205		= NewVectorValue("a_k205", "", "delta_mxy projection onto psi_k205", GetModeAmplitudek205 )
	a_k206		= NewVectorValue("a_k206", "", "delta_mxy projection onto psi_k206", GetModeAmplitudek206 )
	a_k207		= NewVectorValue("a_k207", "", "delta_mxy projection onto psi_k207", GetModeAmplitudek207 )
	a_k208		= NewVectorValue("a_k208", "", "delta_mxy projection onto psi_k208", GetModeAmplitudek208 )
	a_k209		= NewVectorValue("a_k209", "", "delta_mxy projection onto psi_k209", GetModeAmplitudek209 )
	a_k210		= NewVectorValue("a_k210", "", "delta_mxy projection onto psi_k210", GetModeAmplitudek210 )
	a_k211		= NewVectorValue("a_k211", "", "delta_mxy projection onto psi_k211", GetModeAmplitudek211 )
	a_k212		= NewVectorValue("a_k212", "", "delta_mxy projection onto psi_k212", GetModeAmplitudek212 )
	a_k213		= NewVectorValue("a_k213", "", "delta_mxy projection onto psi_k213", GetModeAmplitudek213 )
	a_k214		= NewVectorValue("a_k214", "", "delta_mxy projection onto psi_k214", GetModeAmplitudek214 )
	a_k215		= NewVectorValue("a_k215", "", "delta_mxy projection onto psi_k215", GetModeAmplitudek215 )
	a_k216		= NewVectorValue("a_k216", "", "delta_mxy projection onto psi_k216", GetModeAmplitudek216 )
	a_k217		= NewVectorValue("a_k217", "", "delta_mxy projection onto psi_k217", GetModeAmplitudek217 )
	a_k218		= NewVectorValue("a_k218", "", "delta_mxy projection onto psi_k218", GetModeAmplitudek218 )
	a_k219		= NewVectorValue("a_k219", "", "delta_mxy projection onto psi_k219", GetModeAmplitudek219 )
	a_k220		= NewVectorValue("a_k220", "", "delta_mxy projection onto psi_k220", GetModeAmplitudek220 )
	a_k221		= NewVectorValue("a_k221", "", "delta_mxy projection onto psi_k221", GetModeAmplitudek221 )
	a_k222		= NewVectorValue("a_k222", "", "delta_mxy projection onto psi_k222", GetModeAmplitudek222 )
	a_k223		= NewVectorValue("a_k223", "", "delta_mxy projection onto psi_k223", GetModeAmplitudek223 )
	a_k224		= NewVectorValue("a_k224", "", "delta_mxy projection onto psi_k224", GetModeAmplitudek224 )
	a_k225		= NewVectorValue("a_k225", "", "delta_mxy projection onto psi_k225", GetModeAmplitudek225 )
	a_k226		= NewVectorValue("a_k226", "", "delta_mxy projection onto psi_k226", GetModeAmplitudek226 )
	a_k227		= NewVectorValue("a_k227", "", "delta_mxy projection onto psi_k227", GetModeAmplitudek227 )
	a_k228		= NewVectorValue("a_k228", "", "delta_mxy projection onto psi_k228", GetModeAmplitudek228 )
	a_k229		= NewVectorValue("a_k229", "", "delta_mxy projection onto psi_k229", GetModeAmplitudek229 )
	a_k230		= NewVectorValue("a_k230", "", "delta_mxy projection onto psi_k230", GetModeAmplitudek230 )
	a_k231		= NewVectorValue("a_k231", "", "delta_mxy projection onto psi_k231", GetModeAmplitudek231 )
	a_k232		= NewVectorValue("a_k232", "", "delta_mxy projection onto psi_k232", GetModeAmplitudek232 )
	a_k233		= NewVectorValue("a_k233", "", "delta_mxy projection onto psi_k233", GetModeAmplitudek233 )
	a_k234		= NewVectorValue("a_k234", "", "delta_mxy projection onto psi_k234", GetModeAmplitudek234 )
	a_k235		= NewVectorValue("a_k235", "", "delta_mxy projection onto psi_k235", GetModeAmplitudek235 )
	a_k236		= NewVectorValue("a_k236", "", "delta_mxy projection onto psi_k236", GetModeAmplitudek236 )
	a_k237		= NewVectorValue("a_k237", "", "delta_mxy projection onto psi_k237", GetModeAmplitudek237 )
	a_k238		= NewVectorValue("a_k238", "", "delta_mxy projection onto psi_k238", GetModeAmplitudek238 )
	a_k239		= NewVectorValue("a_k239", "", "delta_mxy projection onto psi_k239", GetModeAmplitudek239 )
	a_k240		= NewVectorValue("a_k240", "", "delta_mxy projection onto psi_k240", GetModeAmplitudek240 )
	a_k241		= NewVectorValue("a_k241", "", "delta_mxy projection onto psi_k241", GetModeAmplitudek241 )
	a_k242		= NewVectorValue("a_k242", "", "delta_mxy projection onto psi_k242", GetModeAmplitudek242 )
	a_k243		= NewVectorValue("a_k243", "", "delta_mxy projection onto psi_k243", GetModeAmplitudek243 )
	a_k244		= NewVectorValue("a_k244", "", "delta_mxy projection onto psi_k244", GetModeAmplitudek244 )
	a_k245		= NewVectorValue("a_k245", "", "delta_mxy projection onto psi_k245", GetModeAmplitudek245 )
	a_k246		= NewVectorValue("a_k246", "", "delta_mxy projection onto psi_k246", GetModeAmplitudek246 )
	a_k247		= NewVectorValue("a_k247", "", "delta_mxy projection onto psi_k247", GetModeAmplitudek247 )
	a_k248		= NewVectorValue("a_k248", "", "delta_mxy projection onto psi_k248", GetModeAmplitudek248 )
	a_k249		= NewVectorValue("a_k249", "", "delta_mxy projection onto psi_k249", GetModeAmplitudek249 )
	a_k250		= NewVectorValue("a_k250", "", "delta_mxy projection onto psi_k250", GetModeAmplitudek250 )
	a_k251		= NewVectorValue("a_k251", "", "delta_mxy projection onto psi_k251", GetModeAmplitudek251 )
	a_k252		= NewVectorValue("a_k252", "", "delta_mxy projection onto psi_k252", GetModeAmplitudek252 )
	a_k253		= NewVectorValue("a_k253", "", "delta_mxy projection onto psi_k253", GetModeAmplitudek253 )
	a_k254		= NewVectorValue("a_k254", "", "delta_mxy projection onto psi_k254", GetModeAmplitudek254 )
	a_k255		= NewVectorValue("a_k255", "", "delta_mxy projection onto psi_k255", GetModeAmplitudek255 )
	a_k256		= NewVectorValue("a_k256", "", "delta_mxy projection onto psi_k256", GetModeAmplitudek256 )
	a_k257		= NewVectorValue("a_k257", "", "delta_mxy projection onto psi_k257", GetModeAmplitudek257 )
	a_k258		= NewVectorValue("a_k258", "", "delta_mxy projection onto psi_k258", GetModeAmplitudek258 )
	a_k259		= NewVectorValue("a_k259", "", "delta_mxy projection onto psi_k259", GetModeAmplitudek259 )
	a_k260		= NewVectorValue("a_k260", "", "delta_mxy projection onto psi_k260", GetModeAmplitudek260 )
	a_k261		= NewVectorValue("a_k261", "", "delta_mxy projection onto psi_k261", GetModeAmplitudek261 )
	a_k262		= NewVectorValue("a_k262", "", "delta_mxy projection onto psi_k262", GetModeAmplitudek262 )
	a_k263		= NewVectorValue("a_k263", "", "delta_mxy projection onto psi_k263", GetModeAmplitudek263 )
	a_k264		= NewVectorValue("a_k264", "", "delta_mxy projection onto psi_k264", GetModeAmplitudek264 )
	a_k265		= NewVectorValue("a_k265", "", "delta_mxy projection onto psi_k265", GetModeAmplitudek265 )
	a_k266		= NewVectorValue("a_k266", "", "delta_mxy projection onto psi_k266", GetModeAmplitudek266 )
	a_k267		= NewVectorValue("a_k267", "", "delta_mxy projection onto psi_k267", GetModeAmplitudek267 )
	a_k268		= NewVectorValue("a_k268", "", "delta_mxy projection onto psi_k268", GetModeAmplitudek268 )
	a_k269		= NewVectorValue("a_k269", "", "delta_mxy projection onto psi_k269", GetModeAmplitudek269 )
	a_k270		= NewVectorValue("a_k270", "", "delta_mxy projection onto psi_k270", GetModeAmplitudek270 )
	a_k271		= NewVectorValue("a_k271", "", "delta_mxy projection onto psi_k271", GetModeAmplitudek271 )
	a_k272		= NewVectorValue("a_k272", "", "delta_mxy projection onto psi_k272", GetModeAmplitudek272 )
	a_k273		= NewVectorValue("a_k273", "", "delta_mxy projection onto psi_k273", GetModeAmplitudek273 )
	a_k274		= NewVectorValue("a_k274", "", "delta_mxy projection onto psi_k274", GetModeAmplitudek274 )
	a_k275		= NewVectorValue("a_k275", "", "delta_mxy projection onto psi_k275", GetModeAmplitudek275 )
	a_k276		= NewVectorValue("a_k276", "", "delta_mxy projection onto psi_k276", GetModeAmplitudek276 )
	a_k277		= NewVectorValue("a_k277", "", "delta_mxy projection onto psi_k277", GetModeAmplitudek277 )
	a_k278		= NewVectorValue("a_k278", "", "delta_mxy projection onto psi_k278", GetModeAmplitudek278 )
	a_k279		= NewVectorValue("a_k279", "", "delta_mxy projection onto psi_k279", GetModeAmplitudek279 )
	a_k280		= NewVectorValue("a_k280", "", "delta_mxy projection onto psi_k280", GetModeAmplitudek280 )
	a_k281		= NewVectorValue("a_k281", "", "delta_mxy projection onto psi_k281", GetModeAmplitudek281 )
	a_k282		= NewVectorValue("a_k282", "", "delta_mxy projection onto psi_k282", GetModeAmplitudek282 )
	a_k283		= NewVectorValue("a_k283", "", "delta_mxy projection onto psi_k283", GetModeAmplitudek283 )
	a_k284		= NewVectorValue("a_k284", "", "delta_mxy projection onto psi_k284", GetModeAmplitudek284 )
	a_k285		= NewVectorValue("a_k285", "", "delta_mxy projection onto psi_k285", GetModeAmplitudek285 )
	a_k286		= NewVectorValue("a_k286", "", "delta_mxy projection onto psi_k286", GetModeAmplitudek286 )
	a_k287		= NewVectorValue("a_k287", "", "delta_mxy projection onto psi_k287", GetModeAmplitudek287 )
	a_k288		= NewVectorValue("a_k288", "", "delta_mxy projection onto psi_k288", GetModeAmplitudek288 )
	a_k289		= NewVectorValue("a_k289", "", "delta_mxy projection onto psi_k289", GetModeAmplitudek289 )
	a_k290		= NewVectorValue("a_k290", "", "delta_mxy projection onto psi_k290", GetModeAmplitudek290 )
	a_k291		= NewVectorValue("a_k291", "", "delta_mxy projection onto psi_k291", GetModeAmplitudek291 )
	a_k292		= NewVectorValue("a_k292", "", "delta_mxy projection onto psi_k292", GetModeAmplitudek292 )
	a_k293		= NewVectorValue("a_k293", "", "delta_mxy projection onto psi_k293", GetModeAmplitudek293 )
	a_k294		= NewVectorValue("a_k294", "", "delta_mxy projection onto psi_k294", GetModeAmplitudek294 )
	a_k295		= NewVectorValue("a_k295", "", "delta_mxy projection onto psi_k295", GetModeAmplitudek295 )
	a_k296		= NewVectorValue("a_k296", "", "delta_mxy projection onto psi_k296", GetModeAmplitudek296 )
	a_k297		= NewVectorValue("a_k297", "", "delta_mxy projection onto psi_k297", GetModeAmplitudek297 )
	a_k298		= NewVectorValue("a_k298", "", "delta_mxy projection onto psi_k298", GetModeAmplitudek298 )
	a_k299		= NewVectorValue("a_k299", "", "delta_mxy projection onto psi_k299", GetModeAmplitudek299 )
	a_k300		= NewVectorValue("a_k300", "", "delta_mxy projection onto psi_k300", GetModeAmplitudek300 )
	a_k301		= NewVectorValue("a_k301", "", "delta_mxy projection onto psi_k301", GetModeAmplitudek301 )
	a_k302		= NewVectorValue("a_k302", "", "delta_mxy projection onto psi_k302", GetModeAmplitudek302 )
	a_k303		= NewVectorValue("a_k303", "", "delta_mxy projection onto psi_k303", GetModeAmplitudek303 )
	a_k304		= NewVectorValue("a_k304", "", "delta_mxy projection onto psi_k304", GetModeAmplitudek304 )
	a_k305		= NewVectorValue("a_k305", "", "delta_mxy projection onto psi_k305", GetModeAmplitudek305 )
	a_k306		= NewVectorValue("a_k306", "", "delta_mxy projection onto psi_k306", GetModeAmplitudek306 )
	a_k307		= NewVectorValue("a_k307", "", "delta_mxy projection onto psi_k307", GetModeAmplitudek307 )
	a_k308		= NewVectorValue("a_k308", "", "delta_mxy projection onto psi_k308", GetModeAmplitudek308 )
	a_k309		= NewVectorValue("a_k309", "", "delta_mxy projection onto psi_k309", GetModeAmplitudek309 )
	a_k310		= NewVectorValue("a_k310", "", "delta_mxy projection onto psi_k310", GetModeAmplitudek310 )
	a_k311		= NewVectorValue("a_k311", "", "delta_mxy projection onto psi_k311", GetModeAmplitudek311 )
	a_k312		= NewVectorValue("a_k312", "", "delta_mxy projection onto psi_k312", GetModeAmplitudek312 )
	a_k313		= NewVectorValue("a_k313", "", "delta_mxy projection onto psi_k313", GetModeAmplitudek313 )
	a_k314		= NewVectorValue("a_k314", "", "delta_mxy projection onto psi_k314", GetModeAmplitudek314 )
	a_k315		= NewVectorValue("a_k315", "", "delta_mxy projection onto psi_k315", GetModeAmplitudek315 )
	a_k316		= NewVectorValue("a_k316", "", "delta_mxy projection onto psi_k316", GetModeAmplitudek316 )
	a_k317		= NewVectorValue("a_k317", "", "delta_mxy projection onto psi_k317", GetModeAmplitudek317 )
	a_k318		= NewVectorValue("a_k318", "", "delta_mxy projection onto psi_k318", GetModeAmplitudek318 )
	a_k319		= NewVectorValue("a_k319", "", "delta_mxy projection onto psi_k319", GetModeAmplitudek319 )
	a_k320		= NewVectorValue("a_k320", "", "delta_mxy projection onto psi_k320", GetModeAmplitudek320 )
	a_k321		= NewVectorValue("a_k321", "", "delta_mxy projection onto psi_k321", GetModeAmplitudek321 )
	a_k322		= NewVectorValue("a_k322", "", "delta_mxy projection onto psi_k322", GetModeAmplitudek322 )
	a_k323		= NewVectorValue("a_k323", "", "delta_mxy projection onto psi_k323", GetModeAmplitudek323 )
	a_k324		= NewVectorValue("a_k324", "", "delta_mxy projection onto psi_k324", GetModeAmplitudek324 )
	a_k325		= NewVectorValue("a_k325", "", "delta_mxy projection onto psi_k325", GetModeAmplitudek325 )
	a_k326		= NewVectorValue("a_k326", "", "delta_mxy projection onto psi_k326", GetModeAmplitudek326 )
	a_k327		= NewVectorValue("a_k327", "", "delta_mxy projection onto psi_k327", GetModeAmplitudek327 )
	a_k328		= NewVectorValue("a_k328", "", "delta_mxy projection onto psi_k328", GetModeAmplitudek328 )
	a_k329		= NewVectorValue("a_k329", "", "delta_mxy projection onto psi_k329", GetModeAmplitudek329 )
	a_k330		= NewVectorValue("a_k330", "", "delta_mxy projection onto psi_k330", GetModeAmplitudek330 )
	a_k331		= NewVectorValue("a_k331", "", "delta_mxy projection onto psi_k331", GetModeAmplitudek331 )
	a_k332		= NewVectorValue("a_k332", "", "delta_mxy projection onto psi_k332", GetModeAmplitudek332 )
	a_k333		= NewVectorValue("a_k333", "", "delta_mxy projection onto psi_k333", GetModeAmplitudek333 )
	a_k334		= NewVectorValue("a_k334", "", "delta_mxy projection onto psi_k334", GetModeAmplitudek334 )
	a_k335		= NewVectorValue("a_k335", "", "delta_mxy projection onto psi_k335", GetModeAmplitudek335 )
	a_k336		= NewVectorValue("a_k336", "", "delta_mxy projection onto psi_k336", GetModeAmplitudek336 )
	a_k337		= NewVectorValue("a_k337", "", "delta_mxy projection onto psi_k337", GetModeAmplitudek337 )
	a_k338		= NewVectorValue("a_k338", "", "delta_mxy projection onto psi_k338", GetModeAmplitudek338 )
	a_k339		= NewVectorValue("a_k339", "", "delta_mxy projection onto psi_k339", GetModeAmplitudek339 )
	a_k340		= NewVectorValue("a_k340", "", "delta_mxy projection onto psi_k340", GetModeAmplitudek340 )
	a_k341		= NewVectorValue("a_k341", "", "delta_mxy projection onto psi_k341", GetModeAmplitudek341 )
	a_k342		= NewVectorValue("a_k342", "", "delta_mxy projection onto psi_k342", GetModeAmplitudek342 )
	a_k343		= NewVectorValue("a_k343", "", "delta_mxy projection onto psi_k343", GetModeAmplitudek343 )
	a_k344		= NewVectorValue("a_k344", "", "delta_mxy projection onto psi_k344", GetModeAmplitudek344 )
	a_k345		= NewVectorValue("a_k345", "", "delta_mxy projection onto psi_k345", GetModeAmplitudek345 )
	a_k346		= NewVectorValue("a_k346", "", "delta_mxy projection onto psi_k346", GetModeAmplitudek346 )
	a_k347		= NewVectorValue("a_k347", "", "delta_mxy projection onto psi_k347", GetModeAmplitudek347 )
	a_k348		= NewVectorValue("a_k348", "", "delta_mxy projection onto psi_k348", GetModeAmplitudek348 )
	a_k349		= NewVectorValue("a_k349", "", "delta_mxy projection onto psi_k349", GetModeAmplitudek349 )
	a_k350		= NewVectorValue("a_k350", "", "delta_mxy projection onto psi_k350", GetModeAmplitudek350 )
	a_k351		= NewVectorValue("a_k351", "", "delta_mxy projection onto psi_k351", GetModeAmplitudek351 )
	a_k352		= NewVectorValue("a_k352", "", "delta_mxy projection onto psi_k352", GetModeAmplitudek352 )
	a_k353		= NewVectorValue("a_k353", "", "delta_mxy projection onto psi_k353", GetModeAmplitudek353 )
	a_k354		= NewVectorValue("a_k354", "", "delta_mxy projection onto psi_k354", GetModeAmplitudek354 )
	a_k355		= NewVectorValue("a_k355", "", "delta_mxy projection onto psi_k355", GetModeAmplitudek355 )
	a_k356		= NewVectorValue("a_k356", "", "delta_mxy projection onto psi_k356", GetModeAmplitudek356 )
	a_k357		= NewVectorValue("a_k357", "", "delta_mxy projection onto psi_k357", GetModeAmplitudek357 )
	a_k358		= NewVectorValue("a_k358", "", "delta_mxy projection onto psi_k358", GetModeAmplitudek358 )
	a_k359		= NewVectorValue("a_k359", "", "delta_mxy projection onto psi_k359", GetModeAmplitudek359 )
	a_k360		= NewVectorValue("a_k360", "", "delta_mxy projection onto psi_k360", GetModeAmplitudek360 )
	a_k361		= NewVectorValue("a_k361", "", "delta_mxy projection onto psi_k361", GetModeAmplitudek361 )
	a_k362		= NewVectorValue("a_k362", "", "delta_mxy projection onto psi_k362", GetModeAmplitudek362 )
	a_k363		= NewVectorValue("a_k363", "", "delta_mxy projection onto psi_k363", GetModeAmplitudek363 )
	a_k364		= NewVectorValue("a_k364", "", "delta_mxy projection onto psi_k364", GetModeAmplitudek364 )
	a_k365		= NewVectorValue("a_k365", "", "delta_mxy projection onto psi_k365", GetModeAmplitudek365 )
	a_k366		= NewVectorValue("a_k366", "", "delta_mxy projection onto psi_k366", GetModeAmplitudek366 )
	a_k367		= NewVectorValue("a_k367", "", "delta_mxy projection onto psi_k367", GetModeAmplitudek367 )
	a_k368		= NewVectorValue("a_k368", "", "delta_mxy projection onto psi_k368", GetModeAmplitudek368 )
	a_k369		= NewVectorValue("a_k369", "", "delta_mxy projection onto psi_k369", GetModeAmplitudek369 )
	a_k370		= NewVectorValue("a_k370", "", "delta_mxy projection onto psi_k370", GetModeAmplitudek370 )
	a_k371		= NewVectorValue("a_k371", "", "delta_mxy projection onto psi_k371", GetModeAmplitudek371 )
	a_k372		= NewVectorValue("a_k372", "", "delta_mxy projection onto psi_k372", GetModeAmplitudek372 )
	a_k373		= NewVectorValue("a_k373", "", "delta_mxy projection onto psi_k373", GetModeAmplitudek373 )
	a_k374		= NewVectorValue("a_k374", "", "delta_mxy projection onto psi_k374", GetModeAmplitudek374 )
	a_k375		= NewVectorValue("a_k375", "", "delta_mxy projection onto psi_k375", GetModeAmplitudek375 )
	a_k376		= NewVectorValue("a_k376", "", "delta_mxy projection onto psi_k376", GetModeAmplitudek376 )
	a_k377		= NewVectorValue("a_k377", "", "delta_mxy projection onto psi_k377", GetModeAmplitudek377 )
	a_k378		= NewVectorValue("a_k378", "", "delta_mxy projection onto psi_k378", GetModeAmplitudek378 )
	a_k379		= NewVectorValue("a_k379", "", "delta_mxy projection onto psi_k379", GetModeAmplitudek379 )
	a_k380		= NewVectorValue("a_k380", "", "delta_mxy projection onto psi_k380", GetModeAmplitudek380 )
	a_k381		= NewVectorValue("a_k381", "", "delta_mxy projection onto psi_k381", GetModeAmplitudek381 )
	a_k382		= NewVectorValue("a_k382", "", "delta_mxy projection onto psi_k382", GetModeAmplitudek382 )
	a_k383		= NewVectorValue("a_k383", "", "delta_mxy projection onto psi_k383", GetModeAmplitudek383 )
	a_k384		= NewVectorValue("a_k384", "", "delta_mxy projection onto psi_k384", GetModeAmplitudek384 )
	a_k385		= NewVectorValue("a_k385", "", "delta_mxy projection onto psi_k385", GetModeAmplitudek385 )
	a_k386		= NewVectorValue("a_k386", "", "delta_mxy projection onto psi_k386", GetModeAmplitudek386 )
	a_k387		= NewVectorValue("a_k387", "", "delta_mxy projection onto psi_k387", GetModeAmplitudek387 )
	a_k388		= NewVectorValue("a_k388", "", "delta_mxy projection onto psi_k388", GetModeAmplitudek388 )
	a_k389		= NewVectorValue("a_k389", "", "delta_mxy projection onto psi_k389", GetModeAmplitudek389 )
	a_k390		= NewVectorValue("a_k390", "", "delta_mxy projection onto psi_k390", GetModeAmplitudek390 )
	a_k391		= NewVectorValue("a_k391", "", "delta_mxy projection onto psi_k391", GetModeAmplitudek391 )
	a_k392		= NewVectorValue("a_k392", "", "delta_mxy projection onto psi_k392", GetModeAmplitudek392 )
	a_k393		= NewVectorValue("a_k393", "", "delta_mxy projection onto psi_k393", GetModeAmplitudek393 )
	a_k394		= NewVectorValue("a_k394", "", "delta_mxy projection onto psi_k394", GetModeAmplitudek394 )
	a_k395		= NewVectorValue("a_k395", "", "delta_mxy projection onto psi_k395", GetModeAmplitudek395 )
	a_k396		= NewVectorValue("a_k396", "", "delta_mxy projection onto psi_k396", GetModeAmplitudek396 )
	a_k397		= NewVectorValue("a_k397", "", "delta_mxy projection onto psi_k397", GetModeAmplitudek397 )
	a_k398		= NewVectorValue("a_k398", "", "delta_mxy projection onto psi_k398", GetModeAmplitudek398 )
	a_k399		= NewVectorValue("a_k399", "", "delta_mxy projection onto psi_k399", GetModeAmplitudek399 )
	a_k400		= NewVectorValue("a_k400", "", "delta_mxy projection onto psi_k400", GetModeAmplitudek400 )
	a_k401		= NewVectorValue("a_k401", "", "delta_mxy projection onto psi_k401", GetModeAmplitudek401 )
	a_k402		= NewVectorValue("a_k402", "", "delta_mxy projection onto psi_k402", GetModeAmplitudek402 )
	a_k403		= NewVectorValue("a_k403", "", "delta_mxy projection onto psi_k403", GetModeAmplitudek403 )
	a_k404		= NewVectorValue("a_k404", "", "delta_mxy projection onto psi_k404", GetModeAmplitudek404 )
	a_k405		= NewVectorValue("a_k405", "", "delta_mxy projection onto psi_k405", GetModeAmplitudek405 )
	a_k406		= NewVectorValue("a_k406", "", "delta_mxy projection onto psi_k406", GetModeAmplitudek406 )
	a_k407		= NewVectorValue("a_k407", "", "delta_mxy projection onto psi_k407", GetModeAmplitudek407 )
	a_k408		= NewVectorValue("a_k408", "", "delta_mxy projection onto psi_k408", GetModeAmplitudek408 )
	a_k409		= NewVectorValue("a_k409", "", "delta_mxy projection onto psi_k409", GetModeAmplitudek409 )
	a_k410		= NewVectorValue("a_k410", "", "delta_mxy projection onto psi_k410", GetModeAmplitudek410 )
	a_k411		= NewVectorValue("a_k411", "", "delta_mxy projection onto psi_k411", GetModeAmplitudek411 )
	a_k412		= NewVectorValue("a_k412", "", "delta_mxy projection onto psi_k412", GetModeAmplitudek412 )
	a_k413		= NewVectorValue("a_k413", "", "delta_mxy projection onto psi_k413", GetModeAmplitudek413 )
	a_k414		= NewVectorValue("a_k414", "", "delta_mxy projection onto psi_k414", GetModeAmplitudek414 )
	a_k415		= NewVectorValue("a_k415", "", "delta_mxy projection onto psi_k415", GetModeAmplitudek415 )
	a_k416		= NewVectorValue("a_k416", "", "delta_mxy projection onto psi_k416", GetModeAmplitudek416 )
	a_k417		= NewVectorValue("a_k417", "", "delta_mxy projection onto psi_k417", GetModeAmplitudek417 )
	a_k418		= NewVectorValue("a_k418", "", "delta_mxy projection onto psi_k418", GetModeAmplitudek418 )
	a_k419		= NewVectorValue("a_k419", "", "delta_mxy projection onto psi_k419", GetModeAmplitudek419 )
	a_k420		= NewVectorValue("a_k420", "", "delta_mxy projection onto psi_k420", GetModeAmplitudek420 )
	a_k421		= NewVectorValue("a_k421", "", "delta_mxy projection onto psi_k421", GetModeAmplitudek421 )
	a_k422		= NewVectorValue("a_k422", "", "delta_mxy projection onto psi_k422", GetModeAmplitudek422 )
	a_k423		= NewVectorValue("a_k423", "", "delta_mxy projection onto psi_k423", GetModeAmplitudek423 )
	a_k424		= NewVectorValue("a_k424", "", "delta_mxy projection onto psi_k424", GetModeAmplitudek424 )
	a_k425		= NewVectorValue("a_k425", "", "delta_mxy projection onto psi_k425", GetModeAmplitudek425 )
	a_k426		= NewVectorValue("a_k426", "", "delta_mxy projection onto psi_k426", GetModeAmplitudek426 )
	a_k427		= NewVectorValue("a_k427", "", "delta_mxy projection onto psi_k427", GetModeAmplitudek427 )
	a_k428		= NewVectorValue("a_k428", "", "delta_mxy projection onto psi_k428", GetModeAmplitudek428 )
	a_k429		= NewVectorValue("a_k429", "", "delta_mxy projection onto psi_k429", GetModeAmplitudek429 )
	a_k430		= NewVectorValue("a_k430", "", "delta_mxy projection onto psi_k430", GetModeAmplitudek430 )
	a_k431		= NewVectorValue("a_k431", "", "delta_mxy projection onto psi_k431", GetModeAmplitudek431 )
	a_k432		= NewVectorValue("a_k432", "", "delta_mxy projection onto psi_k432", GetModeAmplitudek432 )
	a_k433		= NewVectorValue("a_k433", "", "delta_mxy projection onto psi_k433", GetModeAmplitudek433 )
	a_k434		= NewVectorValue("a_k434", "", "delta_mxy projection onto psi_k434", GetModeAmplitudek434 )
	a_k435		= NewVectorValue("a_k435", "", "delta_mxy projection onto psi_k435", GetModeAmplitudek435 )
	a_k436		= NewVectorValue("a_k436", "", "delta_mxy projection onto psi_k436", GetModeAmplitudek436 )
	a_k437		= NewVectorValue("a_k437", "", "delta_mxy projection onto psi_k437", GetModeAmplitudek437 )
	a_k438		= NewVectorValue("a_k438", "", "delta_mxy projection onto psi_k438", GetModeAmplitudek438 )
	a_k439		= NewVectorValue("a_k439", "", "delta_mxy projection onto psi_k439", GetModeAmplitudek439 )
	a_k440		= NewVectorValue("a_k440", "", "delta_mxy projection onto psi_k440", GetModeAmplitudek440 )
	a_k441		= NewVectorValue("a_k441", "", "delta_mxy projection onto psi_k441", GetModeAmplitudek441 )
	a_k442		= NewVectorValue("a_k442", "", "delta_mxy projection onto psi_k442", GetModeAmplitudek442 )
	a_k443		= NewVectorValue("a_k443", "", "delta_mxy projection onto psi_k443", GetModeAmplitudek443 )
	a_k444		= NewVectorValue("a_k444", "", "delta_mxy projection onto psi_k444", GetModeAmplitudek444 )
	a_k445		= NewVectorValue("a_k445", "", "delta_mxy projection onto psi_k445", GetModeAmplitudek445 )
	a_k446		= NewVectorValue("a_k446", "", "delta_mxy projection onto psi_k446", GetModeAmplitudek446 )
	a_k447		= NewVectorValue("a_k447", "", "delta_mxy projection onto psi_k447", GetModeAmplitudek447 )
	a_k448		= NewVectorValue("a_k448", "", "delta_mxy projection onto psi_k448", GetModeAmplitudek448 )
	a_k449		= NewVectorValue("a_k449", "", "delta_mxy projection onto psi_k449", GetModeAmplitudek449 )
	a_k450		= NewVectorValue("a_k450", "", "delta_mxy projection onto psi_k450", GetModeAmplitudek450 )
	a_k451		= NewVectorValue("a_k451", "", "delta_mxy projection onto psi_k451", GetModeAmplitudek451 )
	a_k452		= NewVectorValue("a_k452", "", "delta_mxy projection onto psi_k452", GetModeAmplitudek452 )
	a_k453		= NewVectorValue("a_k453", "", "delta_mxy projection onto psi_k453", GetModeAmplitudek453 )
	a_k454		= NewVectorValue("a_k454", "", "delta_mxy projection onto psi_k454", GetModeAmplitudek454 )
	a_k455		= NewVectorValue("a_k455", "", "delta_mxy projection onto psi_k455", GetModeAmplitudek455 )
	a_k456		= NewVectorValue("a_k456", "", "delta_mxy projection onto psi_k456", GetModeAmplitudek456 )
	a_k457		= NewVectorValue("a_k457", "", "delta_mxy projection onto psi_k457", GetModeAmplitudek457 )
	a_k458		= NewVectorValue("a_k458", "", "delta_mxy projection onto psi_k458", GetModeAmplitudek458 )
	a_k459		= NewVectorValue("a_k459", "", "delta_mxy projection onto psi_k459", GetModeAmplitudek459 )
	a_k460		= NewVectorValue("a_k460", "", "delta_mxy projection onto psi_k460", GetModeAmplitudek460 )
	a_k461		= NewVectorValue("a_k461", "", "delta_mxy projection onto psi_k461", GetModeAmplitudek461 )
	a_k462		= NewVectorValue("a_k462", "", "delta_mxy projection onto psi_k462", GetModeAmplitudek462 )
	a_k463		= NewVectorValue("a_k463", "", "delta_mxy projection onto psi_k463", GetModeAmplitudek463 )
	a_k464		= NewVectorValue("a_k464", "", "delta_mxy projection onto psi_k464", GetModeAmplitudek464 )
	a_k465		= NewVectorValue("a_k465", "", "delta_mxy projection onto psi_k465", GetModeAmplitudek465 )
	a_k466		= NewVectorValue("a_k466", "", "delta_mxy projection onto psi_k466", GetModeAmplitudek466 )
	a_k467		= NewVectorValue("a_k467", "", "delta_mxy projection onto psi_k467", GetModeAmplitudek467 )
	a_k468		= NewVectorValue("a_k468", "", "delta_mxy projection onto psi_k468", GetModeAmplitudek468 )
	a_k469		= NewVectorValue("a_k469", "", "delta_mxy projection onto psi_k469", GetModeAmplitudek469 )
	a_k470		= NewVectorValue("a_k470", "", "delta_mxy projection onto psi_k470", GetModeAmplitudek470 )
	a_k471		= NewVectorValue("a_k471", "", "delta_mxy projection onto psi_k471", GetModeAmplitudek471 )
	a_k472		= NewVectorValue("a_k472", "", "delta_mxy projection onto psi_k472", GetModeAmplitudek472 )
	a_k473		= NewVectorValue("a_k473", "", "delta_mxy projection onto psi_k473", GetModeAmplitudek473 )
	a_k474		= NewVectorValue("a_k474", "", "delta_mxy projection onto psi_k474", GetModeAmplitudek474 )
	a_k475		= NewVectorValue("a_k475", "", "delta_mxy projection onto psi_k475", GetModeAmplitudek475 )
	a_k476		= NewVectorValue("a_k476", "", "delta_mxy projection onto psi_k476", GetModeAmplitudek476 )
	a_k477		= NewVectorValue("a_k477", "", "delta_mxy projection onto psi_k477", GetModeAmplitudek477 )
	a_k478		= NewVectorValue("a_k478", "", "delta_mxy projection onto psi_k478", GetModeAmplitudek478 )
	a_k479		= NewVectorValue("a_k479", "", "delta_mxy projection onto psi_k479", GetModeAmplitudek479 )
	a_k480		= NewVectorValue("a_k480", "", "delta_mxy projection onto psi_k480", GetModeAmplitudek480 )
	a_k481		= NewVectorValue("a_k481", "", "delta_mxy projection onto psi_k481", GetModeAmplitudek481 )
	a_k482		= NewVectorValue("a_k482", "", "delta_mxy projection onto psi_k482", GetModeAmplitudek482 )
	a_k483		= NewVectorValue("a_k483", "", "delta_mxy projection onto psi_k483", GetModeAmplitudek483 )
	a_k484		= NewVectorValue("a_k484", "", "delta_mxy projection onto psi_k484", GetModeAmplitudek484 )
	a_k485		= NewVectorValue("a_k485", "", "delta_mxy projection onto psi_k485", GetModeAmplitudek485 )
	a_k486		= NewVectorValue("a_k486", "", "delta_mxy projection onto psi_k486", GetModeAmplitudek486 )
	a_k487		= NewVectorValue("a_k487", "", "delta_mxy projection onto psi_k487", GetModeAmplitudek487 )
	a_k488		= NewVectorValue("a_k488", "", "delta_mxy projection onto psi_k488", GetModeAmplitudek488 )
	a_k489		= NewVectorValue("a_k489", "", "delta_mxy projection onto psi_k489", GetModeAmplitudek489 )
	a_k490		= NewVectorValue("a_k490", "", "delta_mxy projection onto psi_k490", GetModeAmplitudek490 )
	a_k491		= NewVectorValue("a_k491", "", "delta_mxy projection onto psi_k491", GetModeAmplitudek491 )
	a_k492		= NewVectorValue("a_k492", "", "delta_mxy projection onto psi_k492", GetModeAmplitudek492 )
	a_k493		= NewVectorValue("a_k493", "", "delta_mxy projection onto psi_k493", GetModeAmplitudek493 )
	a_k494		= NewVectorValue("a_k494", "", "delta_mxy projection onto psi_k494", GetModeAmplitudek494 )
	a_k495		= NewVectorValue("a_k495", "", "delta_mxy projection onto psi_k495", GetModeAmplitudek495 )
	a_k496		= NewVectorValue("a_k496", "", "delta_mxy projection onto psi_k496", GetModeAmplitudek496 )
	a_k497		= NewVectorValue("a_k497", "", "delta_mxy projection onto psi_k497", GetModeAmplitudek497 )
	a_k498		= NewVectorValue("a_k498", "", "delta_mxy projection onto psi_k498", GetModeAmplitudek498 )
	a_k499		= NewVectorValue("a_k499", "", "delta_mxy projection onto psi_k499", GetModeAmplitudek499 )
)
	
	
func GetModeAmplitudek000() []float64 {

	sx := Mul(psi_k000, Dot(&M, delta_mx) )
	sy := Mul(psi_k000, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek001() []float64 {

	sx := Mul(psi_k001, Dot(&M, delta_mx) )
	sy := Mul(psi_k001, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek002() []float64 {

	sx := Mul(psi_k002, Dot(&M, delta_mx) )
	sy := Mul(psi_k002, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek003() []float64 {

	sx := Mul(psi_k003, Dot(&M, delta_mx) )
	sy := Mul(psi_k003, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek004() []float64 {

	sx := Mul(psi_k004, Dot(&M, delta_mx) )
	sy := Mul(psi_k004, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek005() []float64 {

	sx := Mul(psi_k005, Dot(&M, delta_mx) )
	sy := Mul(psi_k005, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek006() []float64 {

	sx := Mul(psi_k006, Dot(&M, delta_mx) )
	sy := Mul(psi_k006, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek007() []float64 {

	sx := Mul(psi_k007, Dot(&M, delta_mx) )
	sy := Mul(psi_k007, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek008() []float64 {

	sx := Mul(psi_k008, Dot(&M, delta_mx) )
	sy := Mul(psi_k008, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek009() []float64 {

	sx := Mul(psi_k009, Dot(&M, delta_mx) )
	sy := Mul(psi_k009, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek010() []float64 {

	sx := Mul(psi_k010, Dot(&M, delta_mx) )
	sy := Mul(psi_k010, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek011() []float64 {

	sx := Mul(psi_k011, Dot(&M, delta_mx) )
	sy := Mul(psi_k011, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek012() []float64 {

	sx := Mul(psi_k012, Dot(&M, delta_mx) )
	sy := Mul(psi_k012, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek013() []float64 {

	sx := Mul(psi_k013, Dot(&M, delta_mx) )
	sy := Mul(psi_k013, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek014() []float64 {

	sx := Mul(psi_k014, Dot(&M, delta_mx) )
	sy := Mul(psi_k014, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek015() []float64 {

	sx := Mul(psi_k015, Dot(&M, delta_mx) )
	sy := Mul(psi_k015, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek016() []float64 {

	sx := Mul(psi_k016, Dot(&M, delta_mx) )
	sy := Mul(psi_k016, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek017() []float64 {

	sx := Mul(psi_k017, Dot(&M, delta_mx) )
	sy := Mul(psi_k017, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek018() []float64 {

	sx := Mul(psi_k018, Dot(&M, delta_mx) )
	sy := Mul(psi_k018, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek019() []float64 {

	sx := Mul(psi_k019, Dot(&M, delta_mx) )
	sy := Mul(psi_k019, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek020() []float64 {

	sx := Mul(psi_k020, Dot(&M, delta_mx) )
	sy := Mul(psi_k020, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek021() []float64 {

	sx := Mul(psi_k021, Dot(&M, delta_mx) )
	sy := Mul(psi_k021, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek022() []float64 {

	sx := Mul(psi_k022, Dot(&M, delta_mx) )
	sy := Mul(psi_k022, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek023() []float64 {

	sx := Mul(psi_k023, Dot(&M, delta_mx) )
	sy := Mul(psi_k023, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek024() []float64 {

	sx := Mul(psi_k024, Dot(&M, delta_mx) )
	sy := Mul(psi_k024, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek025() []float64 {

	sx := Mul(psi_k025, Dot(&M, delta_mx) )
	sy := Mul(psi_k025, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek026() []float64 {

	sx := Mul(psi_k026, Dot(&M, delta_mx) )
	sy := Mul(psi_k026, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek027() []float64 {

	sx := Mul(psi_k027, Dot(&M, delta_mx) )
	sy := Mul(psi_k027, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek028() []float64 {

	sx := Mul(psi_k028, Dot(&M, delta_mx) )
	sy := Mul(psi_k028, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek029() []float64 {

	sx := Mul(psi_k029, Dot(&M, delta_mx) )
	sy := Mul(psi_k029, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek030() []float64 {

	sx := Mul(psi_k030, Dot(&M, delta_mx) )
	sy := Mul(psi_k030, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek031() []float64 {

	sx := Mul(psi_k031, Dot(&M, delta_mx) )
	sy := Mul(psi_k031, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek032() []float64 {

	sx := Mul(psi_k032, Dot(&M, delta_mx) )
	sy := Mul(psi_k032, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek033() []float64 {

	sx := Mul(psi_k033, Dot(&M, delta_mx) )
	sy := Mul(psi_k033, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek034() []float64 {

	sx := Mul(psi_k034, Dot(&M, delta_mx) )
	sy := Mul(psi_k034, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek035() []float64 {

	sx := Mul(psi_k035, Dot(&M, delta_mx) )
	sy := Mul(psi_k035, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek036() []float64 {

	sx := Mul(psi_k036, Dot(&M, delta_mx) )
	sy := Mul(psi_k036, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek037() []float64 {

	sx := Mul(psi_k037, Dot(&M, delta_mx) )
	sy := Mul(psi_k037, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek038() []float64 {

	sx := Mul(psi_k038, Dot(&M, delta_mx) )
	sy := Mul(psi_k038, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek039() []float64 {

	sx := Mul(psi_k039, Dot(&M, delta_mx) )
	sy := Mul(psi_k039, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek040() []float64 {

	sx := Mul(psi_k040, Dot(&M, delta_mx) )
	sy := Mul(psi_k040, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek041() []float64 {

	sx := Mul(psi_k041, Dot(&M, delta_mx) )
	sy := Mul(psi_k041, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek042() []float64 {

	sx := Mul(psi_k042, Dot(&M, delta_mx) )
	sy := Mul(psi_k042, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek043() []float64 {

	sx := Mul(psi_k043, Dot(&M, delta_mx) )
	sy := Mul(psi_k043, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek044() []float64 {

	sx := Mul(psi_k044, Dot(&M, delta_mx) )
	sy := Mul(psi_k044, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek045() []float64 {

	sx := Mul(psi_k045, Dot(&M, delta_mx) )
	sy := Mul(psi_k045, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek046() []float64 {

	sx := Mul(psi_k046, Dot(&M, delta_mx) )
	sy := Mul(psi_k046, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek047() []float64 {

	sx := Mul(psi_k047, Dot(&M, delta_mx) )
	sy := Mul(psi_k047, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek048() []float64 {

	sx := Mul(psi_k048, Dot(&M, delta_mx) )
	sy := Mul(psi_k048, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek049() []float64 {

	sx := Mul(psi_k049, Dot(&M, delta_mx) )
	sy := Mul(psi_k049, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek050() []float64 {

	sx := Mul(psi_k050, Dot(&M, delta_mx) )
	sy := Mul(psi_k050, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek051() []float64 {

	sx := Mul(psi_k051, Dot(&M, delta_mx) )
	sy := Mul(psi_k051, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek052() []float64 {

	sx := Mul(psi_k052, Dot(&M, delta_mx) )
	sy := Mul(psi_k052, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek053() []float64 {

	sx := Mul(psi_k053, Dot(&M, delta_mx) )
	sy := Mul(psi_k053, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek054() []float64 {

	sx := Mul(psi_k054, Dot(&M, delta_mx) )
	sy := Mul(psi_k054, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek055() []float64 {

	sx := Mul(psi_k055, Dot(&M, delta_mx) )
	sy := Mul(psi_k055, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek056() []float64 {

	sx := Mul(psi_k056, Dot(&M, delta_mx) )
	sy := Mul(psi_k056, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek057() []float64 {

	sx := Mul(psi_k057, Dot(&M, delta_mx) )
	sy := Mul(psi_k057, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek058() []float64 {

	sx := Mul(psi_k058, Dot(&M, delta_mx) )
	sy := Mul(psi_k058, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek059() []float64 {

	sx := Mul(psi_k059, Dot(&M, delta_mx) )
	sy := Mul(psi_k059, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek060() []float64 {

	sx := Mul(psi_k060, Dot(&M, delta_mx) )
	sy := Mul(psi_k060, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek061() []float64 {

	sx := Mul(psi_k061, Dot(&M, delta_mx) )
	sy := Mul(psi_k061, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek062() []float64 {

	sx := Mul(psi_k062, Dot(&M, delta_mx) )
	sy := Mul(psi_k062, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek063() []float64 {

	sx := Mul(psi_k063, Dot(&M, delta_mx) )
	sy := Mul(psi_k063, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek064() []float64 {

	sx := Mul(psi_k064, Dot(&M, delta_mx) )
	sy := Mul(psi_k064, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek065() []float64 {

	sx := Mul(psi_k065, Dot(&M, delta_mx) )
	sy := Mul(psi_k065, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek066() []float64 {

	sx := Mul(psi_k066, Dot(&M, delta_mx) )
	sy := Mul(psi_k066, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek067() []float64 {

	sx := Mul(psi_k067, Dot(&M, delta_mx) )
	sy := Mul(psi_k067, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek068() []float64 {

	sx := Mul(psi_k068, Dot(&M, delta_mx) )
	sy := Mul(psi_k068, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek069() []float64 {

	sx := Mul(psi_k069, Dot(&M, delta_mx) )
	sy := Mul(psi_k069, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek070() []float64 {

	sx := Mul(psi_k070, Dot(&M, delta_mx) )
	sy := Mul(psi_k070, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek071() []float64 {

	sx := Mul(psi_k071, Dot(&M, delta_mx) )
	sy := Mul(psi_k071, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek072() []float64 {

	sx := Mul(psi_k072, Dot(&M, delta_mx) )
	sy := Mul(psi_k072, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek073() []float64 {

	sx := Mul(psi_k073, Dot(&M, delta_mx) )
	sy := Mul(psi_k073, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek074() []float64 {

	sx := Mul(psi_k074, Dot(&M, delta_mx) )
	sy := Mul(psi_k074, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek075() []float64 {

	sx := Mul(psi_k075, Dot(&M, delta_mx) )
	sy := Mul(psi_k075, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek076() []float64 {

	sx := Mul(psi_k076, Dot(&M, delta_mx) )
	sy := Mul(psi_k076, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek077() []float64 {

	sx := Mul(psi_k077, Dot(&M, delta_mx) )
	sy := Mul(psi_k077, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek078() []float64 {

	sx := Mul(psi_k078, Dot(&M, delta_mx) )
	sy := Mul(psi_k078, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek079() []float64 {

	sx := Mul(psi_k079, Dot(&M, delta_mx) )
	sy := Mul(psi_k079, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek080() []float64 {

	sx := Mul(psi_k080, Dot(&M, delta_mx) )
	sy := Mul(psi_k080, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek081() []float64 {

	sx := Mul(psi_k081, Dot(&M, delta_mx) )
	sy := Mul(psi_k081, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek082() []float64 {

	sx := Mul(psi_k082, Dot(&M, delta_mx) )
	sy := Mul(psi_k082, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek083() []float64 {

	sx := Mul(psi_k083, Dot(&M, delta_mx) )
	sy := Mul(psi_k083, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek084() []float64 {

	sx := Mul(psi_k084, Dot(&M, delta_mx) )
	sy := Mul(psi_k084, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek085() []float64 {

	sx := Mul(psi_k085, Dot(&M, delta_mx) )
	sy := Mul(psi_k085, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek086() []float64 {

	sx := Mul(psi_k086, Dot(&M, delta_mx) )
	sy := Mul(psi_k086, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek087() []float64 {

	sx := Mul(psi_k087, Dot(&M, delta_mx) )
	sy := Mul(psi_k087, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek088() []float64 {

	sx := Mul(psi_k088, Dot(&M, delta_mx) )
	sy := Mul(psi_k088, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek089() []float64 {

	sx := Mul(psi_k089, Dot(&M, delta_mx) )
	sy := Mul(psi_k089, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek090() []float64 {

	sx := Mul(psi_k090, Dot(&M, delta_mx) )
	sy := Mul(psi_k090, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek091() []float64 {

	sx := Mul(psi_k091, Dot(&M, delta_mx) )
	sy := Mul(psi_k091, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek092() []float64 {

	sx := Mul(psi_k092, Dot(&M, delta_mx) )
	sy := Mul(psi_k092, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek093() []float64 {

	sx := Mul(psi_k093, Dot(&M, delta_mx) )
	sy := Mul(psi_k093, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek094() []float64 {

	sx := Mul(psi_k094, Dot(&M, delta_mx) )
	sy := Mul(psi_k094, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek095() []float64 {

	sx := Mul(psi_k095, Dot(&M, delta_mx) )
	sy := Mul(psi_k095, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek096() []float64 {

	sx := Mul(psi_k096, Dot(&M, delta_mx) )
	sy := Mul(psi_k096, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek097() []float64 {

	sx := Mul(psi_k097, Dot(&M, delta_mx) )
	sy := Mul(psi_k097, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek098() []float64 {

	sx := Mul(psi_k098, Dot(&M, delta_mx) )
	sy := Mul(psi_k098, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek099() []float64 {

	sx := Mul(psi_k099, Dot(&M, delta_mx) )
	sy := Mul(psi_k099, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek100() []float64 {

	sx := Mul(psi_k100, Dot(&M, delta_mx) )
	sy := Mul(psi_k100, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek101() []float64 {

	sx := Mul(psi_k101, Dot(&M, delta_mx) )
	sy := Mul(psi_k101, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek102() []float64 {

	sx := Mul(psi_k102, Dot(&M, delta_mx) )
	sy := Mul(psi_k102, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek103() []float64 {

	sx := Mul(psi_k103, Dot(&M, delta_mx) )
	sy := Mul(psi_k103, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek104() []float64 {

	sx := Mul(psi_k104, Dot(&M, delta_mx) )
	sy := Mul(psi_k104, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek105() []float64 {

	sx := Mul(psi_k105, Dot(&M, delta_mx) )
	sy := Mul(psi_k105, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek106() []float64 {

	sx := Mul(psi_k106, Dot(&M, delta_mx) )
	sy := Mul(psi_k106, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek107() []float64 {

	sx := Mul(psi_k107, Dot(&M, delta_mx) )
	sy := Mul(psi_k107, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek108() []float64 {

	sx := Mul(psi_k108, Dot(&M, delta_mx) )
	sy := Mul(psi_k108, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek109() []float64 {

	sx := Mul(psi_k109, Dot(&M, delta_mx) )
	sy := Mul(psi_k109, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek110() []float64 {

	sx := Mul(psi_k110, Dot(&M, delta_mx) )
	sy := Mul(psi_k110, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek111() []float64 {

	sx := Mul(psi_k111, Dot(&M, delta_mx) )
	sy := Mul(psi_k111, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek112() []float64 {

	sx := Mul(psi_k112, Dot(&M, delta_mx) )
	sy := Mul(psi_k112, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek113() []float64 {

	sx := Mul(psi_k113, Dot(&M, delta_mx) )
	sy := Mul(psi_k113, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek114() []float64 {

	sx := Mul(psi_k114, Dot(&M, delta_mx) )
	sy := Mul(psi_k114, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek115() []float64 {

	sx := Mul(psi_k115, Dot(&M, delta_mx) )
	sy := Mul(psi_k115, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek116() []float64 {

	sx := Mul(psi_k116, Dot(&M, delta_mx) )
	sy := Mul(psi_k116, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek117() []float64 {

	sx := Mul(psi_k117, Dot(&M, delta_mx) )
	sy := Mul(psi_k117, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek118() []float64 {

	sx := Mul(psi_k118, Dot(&M, delta_mx) )
	sy := Mul(psi_k118, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek119() []float64 {

	sx := Mul(psi_k119, Dot(&M, delta_mx) )
	sy := Mul(psi_k119, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek120() []float64 {

	sx := Mul(psi_k120, Dot(&M, delta_mx) )
	sy := Mul(psi_k120, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek121() []float64 {

	sx := Mul(psi_k121, Dot(&M, delta_mx) )
	sy := Mul(psi_k121, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek122() []float64 {

	sx := Mul(psi_k122, Dot(&M, delta_mx) )
	sy := Mul(psi_k122, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek123() []float64 {

	sx := Mul(psi_k123, Dot(&M, delta_mx) )
	sy := Mul(psi_k123, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek124() []float64 {

	sx := Mul(psi_k124, Dot(&M, delta_mx) )
	sy := Mul(psi_k124, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek125() []float64 {

	sx := Mul(psi_k125, Dot(&M, delta_mx) )
	sy := Mul(psi_k125, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek126() []float64 {

	sx := Mul(psi_k126, Dot(&M, delta_mx) )
	sy := Mul(psi_k126, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek127() []float64 {

	sx := Mul(psi_k127, Dot(&M, delta_mx) )
	sy := Mul(psi_k127, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek128() []float64 {

	sx := Mul(psi_k128, Dot(&M, delta_mx) )
	sy := Mul(psi_k128, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek129() []float64 {

	sx := Mul(psi_k129, Dot(&M, delta_mx) )
	sy := Mul(psi_k129, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek130() []float64 {

	sx := Mul(psi_k130, Dot(&M, delta_mx) )
	sy := Mul(psi_k130, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek131() []float64 {

	sx := Mul(psi_k131, Dot(&M, delta_mx) )
	sy := Mul(psi_k131, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek132() []float64 {

	sx := Mul(psi_k132, Dot(&M, delta_mx) )
	sy := Mul(psi_k132, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek133() []float64 {

	sx := Mul(psi_k133, Dot(&M, delta_mx) )
	sy := Mul(psi_k133, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek134() []float64 {

	sx := Mul(psi_k134, Dot(&M, delta_mx) )
	sy := Mul(psi_k134, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek135() []float64 {

	sx := Mul(psi_k135, Dot(&M, delta_mx) )
	sy := Mul(psi_k135, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek136() []float64 {

	sx := Mul(psi_k136, Dot(&M, delta_mx) )
	sy := Mul(psi_k136, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek137() []float64 {

	sx := Mul(psi_k137, Dot(&M, delta_mx) )
	sy := Mul(psi_k137, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek138() []float64 {

	sx := Mul(psi_k138, Dot(&M, delta_mx) )
	sy := Mul(psi_k138, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek139() []float64 {

	sx := Mul(psi_k139, Dot(&M, delta_mx) )
	sy := Mul(psi_k139, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek140() []float64 {

	sx := Mul(psi_k140, Dot(&M, delta_mx) )
	sy := Mul(psi_k140, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek141() []float64 {

	sx := Mul(psi_k141, Dot(&M, delta_mx) )
	sy := Mul(psi_k141, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek142() []float64 {

	sx := Mul(psi_k142, Dot(&M, delta_mx) )
	sy := Mul(psi_k142, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek143() []float64 {

	sx := Mul(psi_k143, Dot(&M, delta_mx) )
	sy := Mul(psi_k143, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek144() []float64 {

	sx := Mul(psi_k144, Dot(&M, delta_mx) )
	sy := Mul(psi_k144, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek145() []float64 {

	sx := Mul(psi_k145, Dot(&M, delta_mx) )
	sy := Mul(psi_k145, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek146() []float64 {

	sx := Mul(psi_k146, Dot(&M, delta_mx) )
	sy := Mul(psi_k146, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek147() []float64 {

	sx := Mul(psi_k147, Dot(&M, delta_mx) )
	sy := Mul(psi_k147, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek148() []float64 {

	sx := Mul(psi_k148, Dot(&M, delta_mx) )
	sy := Mul(psi_k148, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek149() []float64 {

	sx := Mul(psi_k149, Dot(&M, delta_mx) )
	sy := Mul(psi_k149, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek150() []float64 {

	sx := Mul(psi_k150, Dot(&M, delta_mx) )
	sy := Mul(psi_k150, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek151() []float64 {

	sx := Mul(psi_k151, Dot(&M, delta_mx) )
	sy := Mul(psi_k151, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek152() []float64 {

	sx := Mul(psi_k152, Dot(&M, delta_mx) )
	sy := Mul(psi_k152, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek153() []float64 {

	sx := Mul(psi_k153, Dot(&M, delta_mx) )
	sy := Mul(psi_k153, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek154() []float64 {

	sx := Mul(psi_k154, Dot(&M, delta_mx) )
	sy := Mul(psi_k154, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek155() []float64 {

	sx := Mul(psi_k155, Dot(&M, delta_mx) )
	sy := Mul(psi_k155, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek156() []float64 {

	sx := Mul(psi_k156, Dot(&M, delta_mx) )
	sy := Mul(psi_k156, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek157() []float64 {

	sx := Mul(psi_k157, Dot(&M, delta_mx) )
	sy := Mul(psi_k157, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek158() []float64 {

	sx := Mul(psi_k158, Dot(&M, delta_mx) )
	sy := Mul(psi_k158, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek159() []float64 {

	sx := Mul(psi_k159, Dot(&M, delta_mx) )
	sy := Mul(psi_k159, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek160() []float64 {

	sx := Mul(psi_k160, Dot(&M, delta_mx) )
	sy := Mul(psi_k160, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek161() []float64 {

	sx := Mul(psi_k161, Dot(&M, delta_mx) )
	sy := Mul(psi_k161, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek162() []float64 {

	sx := Mul(psi_k162, Dot(&M, delta_mx) )
	sy := Mul(psi_k162, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek163() []float64 {

	sx := Mul(psi_k163, Dot(&M, delta_mx) )
	sy := Mul(psi_k163, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek164() []float64 {

	sx := Mul(psi_k164, Dot(&M, delta_mx) )
	sy := Mul(psi_k164, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek165() []float64 {

	sx := Mul(psi_k165, Dot(&M, delta_mx) )
	sy := Mul(psi_k165, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek166() []float64 {

	sx := Mul(psi_k166, Dot(&M, delta_mx) )
	sy := Mul(psi_k166, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek167() []float64 {

	sx := Mul(psi_k167, Dot(&M, delta_mx) )
	sy := Mul(psi_k167, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek168() []float64 {

	sx := Mul(psi_k168, Dot(&M, delta_mx) )
	sy := Mul(psi_k168, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek169() []float64 {

	sx := Mul(psi_k169, Dot(&M, delta_mx) )
	sy := Mul(psi_k169, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek170() []float64 {

	sx := Mul(psi_k170, Dot(&M, delta_mx) )
	sy := Mul(psi_k170, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek171() []float64 {

	sx := Mul(psi_k171, Dot(&M, delta_mx) )
	sy := Mul(psi_k171, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek172() []float64 {

	sx := Mul(psi_k172, Dot(&M, delta_mx) )
	sy := Mul(psi_k172, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek173() []float64 {

	sx := Mul(psi_k173, Dot(&M, delta_mx) )
	sy := Mul(psi_k173, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek174() []float64 {

	sx := Mul(psi_k174, Dot(&M, delta_mx) )
	sy := Mul(psi_k174, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek175() []float64 {

	sx := Mul(psi_k175, Dot(&M, delta_mx) )
	sy := Mul(psi_k175, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek176() []float64 {

	sx := Mul(psi_k176, Dot(&M, delta_mx) )
	sy := Mul(psi_k176, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek177() []float64 {

	sx := Mul(psi_k177, Dot(&M, delta_mx) )
	sy := Mul(psi_k177, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek178() []float64 {

	sx := Mul(psi_k178, Dot(&M, delta_mx) )
	sy := Mul(psi_k178, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek179() []float64 {

	sx := Mul(psi_k179, Dot(&M, delta_mx) )
	sy := Mul(psi_k179, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek180() []float64 {

	sx := Mul(psi_k180, Dot(&M, delta_mx) )
	sy := Mul(psi_k180, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek181() []float64 {

	sx := Mul(psi_k181, Dot(&M, delta_mx) )
	sy := Mul(psi_k181, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek182() []float64 {

	sx := Mul(psi_k182, Dot(&M, delta_mx) )
	sy := Mul(psi_k182, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek183() []float64 {

	sx := Mul(psi_k183, Dot(&M, delta_mx) )
	sy := Mul(psi_k183, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek184() []float64 {

	sx := Mul(psi_k184, Dot(&M, delta_mx) )
	sy := Mul(psi_k184, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek185() []float64 {

	sx := Mul(psi_k185, Dot(&M, delta_mx) )
	sy := Mul(psi_k185, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek186() []float64 {

	sx := Mul(psi_k186, Dot(&M, delta_mx) )
	sy := Mul(psi_k186, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek187() []float64 {

	sx := Mul(psi_k187, Dot(&M, delta_mx) )
	sy := Mul(psi_k187, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek188() []float64 {

	sx := Mul(psi_k188, Dot(&M, delta_mx) )
	sy := Mul(psi_k188, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek189() []float64 {

	sx := Mul(psi_k189, Dot(&M, delta_mx) )
	sy := Mul(psi_k189, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek190() []float64 {

	sx := Mul(psi_k190, Dot(&M, delta_mx) )
	sy := Mul(psi_k190, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek191() []float64 {

	sx := Mul(psi_k191, Dot(&M, delta_mx) )
	sy := Mul(psi_k191, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek192() []float64 {

	sx := Mul(psi_k192, Dot(&M, delta_mx) )
	sy := Mul(psi_k192, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek193() []float64 {

	sx := Mul(psi_k193, Dot(&M, delta_mx) )
	sy := Mul(psi_k193, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek194() []float64 {

	sx := Mul(psi_k194, Dot(&M, delta_mx) )
	sy := Mul(psi_k194, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek195() []float64 {

	sx := Mul(psi_k195, Dot(&M, delta_mx) )
	sy := Mul(psi_k195, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek196() []float64 {

	sx := Mul(psi_k196, Dot(&M, delta_mx) )
	sy := Mul(psi_k196, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek197() []float64 {

	sx := Mul(psi_k197, Dot(&M, delta_mx) )
	sy := Mul(psi_k197, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek198() []float64 {

	sx := Mul(psi_k198, Dot(&M, delta_mx) )
	sy := Mul(psi_k198, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek199() []float64 {

	sx := Mul(psi_k199, Dot(&M, delta_mx) )
	sy := Mul(psi_k199, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek200() []float64 {

	sx := Mul(psi_k200, Dot(&M, delta_mx) )
	sy := Mul(psi_k200, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek201() []float64 {

	sx := Mul(psi_k201, Dot(&M, delta_mx) )
	sy := Mul(psi_k201, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek202() []float64 {

	sx := Mul(psi_k202, Dot(&M, delta_mx) )
	sy := Mul(psi_k202, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek203() []float64 {

	sx := Mul(psi_k203, Dot(&M, delta_mx) )
	sy := Mul(psi_k203, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek204() []float64 {

	sx := Mul(psi_k204, Dot(&M, delta_mx) )
	sy := Mul(psi_k204, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek205() []float64 {

	sx := Mul(psi_k205, Dot(&M, delta_mx) )
	sy := Mul(psi_k205, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek206() []float64 {

	sx := Mul(psi_k206, Dot(&M, delta_mx) )
	sy := Mul(psi_k206, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek207() []float64 {

	sx := Mul(psi_k207, Dot(&M, delta_mx) )
	sy := Mul(psi_k207, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek208() []float64 {

	sx := Mul(psi_k208, Dot(&M, delta_mx) )
	sy := Mul(psi_k208, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek209() []float64 {

	sx := Mul(psi_k209, Dot(&M, delta_mx) )
	sy := Mul(psi_k209, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek210() []float64 {

	sx := Mul(psi_k210, Dot(&M, delta_mx) )
	sy := Mul(psi_k210, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek211() []float64 {

	sx := Mul(psi_k211, Dot(&M, delta_mx) )
	sy := Mul(psi_k211, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek212() []float64 {

	sx := Mul(psi_k212, Dot(&M, delta_mx) )
	sy := Mul(psi_k212, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek213() []float64 {

	sx := Mul(psi_k213, Dot(&M, delta_mx) )
	sy := Mul(psi_k213, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek214() []float64 {

	sx := Mul(psi_k214, Dot(&M, delta_mx) )
	sy := Mul(psi_k214, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek215() []float64 {

	sx := Mul(psi_k215, Dot(&M, delta_mx) )
	sy := Mul(psi_k215, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek216() []float64 {

	sx := Mul(psi_k216, Dot(&M, delta_mx) )
	sy := Mul(psi_k216, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek217() []float64 {

	sx := Mul(psi_k217, Dot(&M, delta_mx) )
	sy := Mul(psi_k217, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek218() []float64 {

	sx := Mul(psi_k218, Dot(&M, delta_mx) )
	sy := Mul(psi_k218, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek219() []float64 {

	sx := Mul(psi_k219, Dot(&M, delta_mx) )
	sy := Mul(psi_k219, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek220() []float64 {

	sx := Mul(psi_k220, Dot(&M, delta_mx) )
	sy := Mul(psi_k220, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek221() []float64 {

	sx := Mul(psi_k221, Dot(&M, delta_mx) )
	sy := Mul(psi_k221, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek222() []float64 {

	sx := Mul(psi_k222, Dot(&M, delta_mx) )
	sy := Mul(psi_k222, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek223() []float64 {

	sx := Mul(psi_k223, Dot(&M, delta_mx) )
	sy := Mul(psi_k223, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek224() []float64 {

	sx := Mul(psi_k224, Dot(&M, delta_mx) )
	sy := Mul(psi_k224, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek225() []float64 {

	sx := Mul(psi_k225, Dot(&M, delta_mx) )
	sy := Mul(psi_k225, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek226() []float64 {

	sx := Mul(psi_k226, Dot(&M, delta_mx) )
	sy := Mul(psi_k226, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek227() []float64 {

	sx := Mul(psi_k227, Dot(&M, delta_mx) )
	sy := Mul(psi_k227, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek228() []float64 {

	sx := Mul(psi_k228, Dot(&M, delta_mx) )
	sy := Mul(psi_k228, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek229() []float64 {

	sx := Mul(psi_k229, Dot(&M, delta_mx) )
	sy := Mul(psi_k229, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek230() []float64 {

	sx := Mul(psi_k230, Dot(&M, delta_mx) )
	sy := Mul(psi_k230, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek231() []float64 {

	sx := Mul(psi_k231, Dot(&M, delta_mx) )
	sy := Mul(psi_k231, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek232() []float64 {

	sx := Mul(psi_k232, Dot(&M, delta_mx) )
	sy := Mul(psi_k232, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek233() []float64 {

	sx := Mul(psi_k233, Dot(&M, delta_mx) )
	sy := Mul(psi_k233, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek234() []float64 {

	sx := Mul(psi_k234, Dot(&M, delta_mx) )
	sy := Mul(psi_k234, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek235() []float64 {

	sx := Mul(psi_k235, Dot(&M, delta_mx) )
	sy := Mul(psi_k235, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek236() []float64 {

	sx := Mul(psi_k236, Dot(&M, delta_mx) )
	sy := Mul(psi_k236, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek237() []float64 {

	sx := Mul(psi_k237, Dot(&M, delta_mx) )
	sy := Mul(psi_k237, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek238() []float64 {

	sx := Mul(psi_k238, Dot(&M, delta_mx) )
	sy := Mul(psi_k238, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek239() []float64 {

	sx := Mul(psi_k239, Dot(&M, delta_mx) )
	sy := Mul(psi_k239, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek240() []float64 {

	sx := Mul(psi_k240, Dot(&M, delta_mx) )
	sy := Mul(psi_k240, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek241() []float64 {

	sx := Mul(psi_k241, Dot(&M, delta_mx) )
	sy := Mul(psi_k241, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek242() []float64 {

	sx := Mul(psi_k242, Dot(&M, delta_mx) )
	sy := Mul(psi_k242, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek243() []float64 {

	sx := Mul(psi_k243, Dot(&M, delta_mx) )
	sy := Mul(psi_k243, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek244() []float64 {

	sx := Mul(psi_k244, Dot(&M, delta_mx) )
	sy := Mul(psi_k244, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek245() []float64 {

	sx := Mul(psi_k245, Dot(&M, delta_mx) )
	sy := Mul(psi_k245, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek246() []float64 {

	sx := Mul(psi_k246, Dot(&M, delta_mx) )
	sy := Mul(psi_k246, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek247() []float64 {

	sx := Mul(psi_k247, Dot(&M, delta_mx) )
	sy := Mul(psi_k247, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek248() []float64 {

	sx := Mul(psi_k248, Dot(&M, delta_mx) )
	sy := Mul(psi_k248, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek249() []float64 {

	sx := Mul(psi_k249, Dot(&M, delta_mx) )
	sy := Mul(psi_k249, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek250() []float64 {

	sx := Mul(psi_k250, Dot(&M, delta_mx) )
	sy := Mul(psi_k250, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek251() []float64 {

	sx := Mul(psi_k251, Dot(&M, delta_mx) )
	sy := Mul(psi_k251, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek252() []float64 {

	sx := Mul(psi_k252, Dot(&M, delta_mx) )
	sy := Mul(psi_k252, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek253() []float64 {

	sx := Mul(psi_k253, Dot(&M, delta_mx) )
	sy := Mul(psi_k253, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek254() []float64 {

	sx := Mul(psi_k254, Dot(&M, delta_mx) )
	sy := Mul(psi_k254, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek255() []float64 {

	sx := Mul(psi_k255, Dot(&M, delta_mx) )
	sy := Mul(psi_k255, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek256() []float64 {

	sx := Mul(psi_k256, Dot(&M, delta_mx) )
	sy := Mul(psi_k256, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek257() []float64 {

	sx := Mul(psi_k257, Dot(&M, delta_mx) )
	sy := Mul(psi_k257, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek258() []float64 {

	sx := Mul(psi_k258, Dot(&M, delta_mx) )
	sy := Mul(psi_k258, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek259() []float64 {

	sx := Mul(psi_k259, Dot(&M, delta_mx) )
	sy := Mul(psi_k259, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek260() []float64 {

	sx := Mul(psi_k260, Dot(&M, delta_mx) )
	sy := Mul(psi_k260, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek261() []float64 {

	sx := Mul(psi_k261, Dot(&M, delta_mx) )
	sy := Mul(psi_k261, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek262() []float64 {

	sx := Mul(psi_k262, Dot(&M, delta_mx) )
	sy := Mul(psi_k262, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek263() []float64 {

	sx := Mul(psi_k263, Dot(&M, delta_mx) )
	sy := Mul(psi_k263, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek264() []float64 {

	sx := Mul(psi_k264, Dot(&M, delta_mx) )
	sy := Mul(psi_k264, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek265() []float64 {

	sx := Mul(psi_k265, Dot(&M, delta_mx) )
	sy := Mul(psi_k265, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek266() []float64 {

	sx := Mul(psi_k266, Dot(&M, delta_mx) )
	sy := Mul(psi_k266, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek267() []float64 {

	sx := Mul(psi_k267, Dot(&M, delta_mx) )
	sy := Mul(psi_k267, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek268() []float64 {

	sx := Mul(psi_k268, Dot(&M, delta_mx) )
	sy := Mul(psi_k268, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek269() []float64 {

	sx := Mul(psi_k269, Dot(&M, delta_mx) )
	sy := Mul(psi_k269, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek270() []float64 {

	sx := Mul(psi_k270, Dot(&M, delta_mx) )
	sy := Mul(psi_k270, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek271() []float64 {

	sx := Mul(psi_k271, Dot(&M, delta_mx) )
	sy := Mul(psi_k271, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek272() []float64 {

	sx := Mul(psi_k272, Dot(&M, delta_mx) )
	sy := Mul(psi_k272, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek273() []float64 {

	sx := Mul(psi_k273, Dot(&M, delta_mx) )
	sy := Mul(psi_k273, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek274() []float64 {

	sx := Mul(psi_k274, Dot(&M, delta_mx) )
	sy := Mul(psi_k274, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek275() []float64 {

	sx := Mul(psi_k275, Dot(&M, delta_mx) )
	sy := Mul(psi_k275, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek276() []float64 {

	sx := Mul(psi_k276, Dot(&M, delta_mx) )
	sy := Mul(psi_k276, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek277() []float64 {

	sx := Mul(psi_k277, Dot(&M, delta_mx) )
	sy := Mul(psi_k277, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek278() []float64 {

	sx := Mul(psi_k278, Dot(&M, delta_mx) )
	sy := Mul(psi_k278, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek279() []float64 {

	sx := Mul(psi_k279, Dot(&M, delta_mx) )
	sy := Mul(psi_k279, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek280() []float64 {

	sx := Mul(psi_k280, Dot(&M, delta_mx) )
	sy := Mul(psi_k280, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek281() []float64 {

	sx := Mul(psi_k281, Dot(&M, delta_mx) )
	sy := Mul(psi_k281, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek282() []float64 {

	sx := Mul(psi_k282, Dot(&M, delta_mx) )
	sy := Mul(psi_k282, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek283() []float64 {

	sx := Mul(psi_k283, Dot(&M, delta_mx) )
	sy := Mul(psi_k283, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek284() []float64 {

	sx := Mul(psi_k284, Dot(&M, delta_mx) )
	sy := Mul(psi_k284, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek285() []float64 {

	sx := Mul(psi_k285, Dot(&M, delta_mx) )
	sy := Mul(psi_k285, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek286() []float64 {

	sx := Mul(psi_k286, Dot(&M, delta_mx) )
	sy := Mul(psi_k286, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek287() []float64 {

	sx := Mul(psi_k287, Dot(&M, delta_mx) )
	sy := Mul(psi_k287, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek288() []float64 {

	sx := Mul(psi_k288, Dot(&M, delta_mx) )
	sy := Mul(psi_k288, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek289() []float64 {

	sx := Mul(psi_k289, Dot(&M, delta_mx) )
	sy := Mul(psi_k289, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek290() []float64 {

	sx := Mul(psi_k290, Dot(&M, delta_mx) )
	sy := Mul(psi_k290, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek291() []float64 {

	sx := Mul(psi_k291, Dot(&M, delta_mx) )
	sy := Mul(psi_k291, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek292() []float64 {

	sx := Mul(psi_k292, Dot(&M, delta_mx) )
	sy := Mul(psi_k292, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek293() []float64 {

	sx := Mul(psi_k293, Dot(&M, delta_mx) )
	sy := Mul(psi_k293, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek294() []float64 {

	sx := Mul(psi_k294, Dot(&M, delta_mx) )
	sy := Mul(psi_k294, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek295() []float64 {

	sx := Mul(psi_k295, Dot(&M, delta_mx) )
	sy := Mul(psi_k295, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek296() []float64 {

	sx := Mul(psi_k296, Dot(&M, delta_mx) )
	sy := Mul(psi_k296, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek297() []float64 {

	sx := Mul(psi_k297, Dot(&M, delta_mx) )
	sy := Mul(psi_k297, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek298() []float64 {

	sx := Mul(psi_k298, Dot(&M, delta_mx) )
	sy := Mul(psi_k298, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek299() []float64 {

	sx := Mul(psi_k299, Dot(&M, delta_mx) )
	sy := Mul(psi_k299, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek300() []float64 {

	sx := Mul(psi_k300, Dot(&M, delta_mx) )
	sy := Mul(psi_k300, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek301() []float64 {

	sx := Mul(psi_k301, Dot(&M, delta_mx) )
	sy := Mul(psi_k301, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek302() []float64 {

	sx := Mul(psi_k302, Dot(&M, delta_mx) )
	sy := Mul(psi_k302, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek303() []float64 {

	sx := Mul(psi_k303, Dot(&M, delta_mx) )
	sy := Mul(psi_k303, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek304() []float64 {

	sx := Mul(psi_k304, Dot(&M, delta_mx) )
	sy := Mul(psi_k304, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek305() []float64 {

	sx := Mul(psi_k305, Dot(&M, delta_mx) )
	sy := Mul(psi_k305, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek306() []float64 {

	sx := Mul(psi_k306, Dot(&M, delta_mx) )
	sy := Mul(psi_k306, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek307() []float64 {

	sx := Mul(psi_k307, Dot(&M, delta_mx) )
	sy := Mul(psi_k307, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek308() []float64 {

	sx := Mul(psi_k308, Dot(&M, delta_mx) )
	sy := Mul(psi_k308, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek309() []float64 {

	sx := Mul(psi_k309, Dot(&M, delta_mx) )
	sy := Mul(psi_k309, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek310() []float64 {

	sx := Mul(psi_k310, Dot(&M, delta_mx) )
	sy := Mul(psi_k310, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek311() []float64 {

	sx := Mul(psi_k311, Dot(&M, delta_mx) )
	sy := Mul(psi_k311, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek312() []float64 {

	sx := Mul(psi_k312, Dot(&M, delta_mx) )
	sy := Mul(psi_k312, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek313() []float64 {

	sx := Mul(psi_k313, Dot(&M, delta_mx) )
	sy := Mul(psi_k313, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek314() []float64 {

	sx := Mul(psi_k314, Dot(&M, delta_mx) )
	sy := Mul(psi_k314, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek315() []float64 {

	sx := Mul(psi_k315, Dot(&M, delta_mx) )
	sy := Mul(psi_k315, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek316() []float64 {

	sx := Mul(psi_k316, Dot(&M, delta_mx) )
	sy := Mul(psi_k316, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek317() []float64 {

	sx := Mul(psi_k317, Dot(&M, delta_mx) )
	sy := Mul(psi_k317, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek318() []float64 {

	sx := Mul(psi_k318, Dot(&M, delta_mx) )
	sy := Mul(psi_k318, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek319() []float64 {

	sx := Mul(psi_k319, Dot(&M, delta_mx) )
	sy := Mul(psi_k319, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek320() []float64 {

	sx := Mul(psi_k320, Dot(&M, delta_mx) )
	sy := Mul(psi_k320, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek321() []float64 {

	sx := Mul(psi_k321, Dot(&M, delta_mx) )
	sy := Mul(psi_k321, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek322() []float64 {

	sx := Mul(psi_k322, Dot(&M, delta_mx) )
	sy := Mul(psi_k322, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek323() []float64 {

	sx := Mul(psi_k323, Dot(&M, delta_mx) )
	sy := Mul(psi_k323, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek324() []float64 {

	sx := Mul(psi_k324, Dot(&M, delta_mx) )
	sy := Mul(psi_k324, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek325() []float64 {

	sx := Mul(psi_k325, Dot(&M, delta_mx) )
	sy := Mul(psi_k325, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek326() []float64 {

	sx := Mul(psi_k326, Dot(&M, delta_mx) )
	sy := Mul(psi_k326, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek327() []float64 {

	sx := Mul(psi_k327, Dot(&M, delta_mx) )
	sy := Mul(psi_k327, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek328() []float64 {

	sx := Mul(psi_k328, Dot(&M, delta_mx) )
	sy := Mul(psi_k328, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek329() []float64 {

	sx := Mul(psi_k329, Dot(&M, delta_mx) )
	sy := Mul(psi_k329, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek330() []float64 {

	sx := Mul(psi_k330, Dot(&M, delta_mx) )
	sy := Mul(psi_k330, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek331() []float64 {

	sx := Mul(psi_k331, Dot(&M, delta_mx) )
	sy := Mul(psi_k331, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek332() []float64 {

	sx := Mul(psi_k332, Dot(&M, delta_mx) )
	sy := Mul(psi_k332, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek333() []float64 {

	sx := Mul(psi_k333, Dot(&M, delta_mx) )
	sy := Mul(psi_k333, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek334() []float64 {

	sx := Mul(psi_k334, Dot(&M, delta_mx) )
	sy := Mul(psi_k334, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek335() []float64 {

	sx := Mul(psi_k335, Dot(&M, delta_mx) )
	sy := Mul(psi_k335, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek336() []float64 {

	sx := Mul(psi_k336, Dot(&M, delta_mx) )
	sy := Mul(psi_k336, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek337() []float64 {

	sx := Mul(psi_k337, Dot(&M, delta_mx) )
	sy := Mul(psi_k337, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek338() []float64 {

	sx := Mul(psi_k338, Dot(&M, delta_mx) )
	sy := Mul(psi_k338, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek339() []float64 {

	sx := Mul(psi_k339, Dot(&M, delta_mx) )
	sy := Mul(psi_k339, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek340() []float64 {

	sx := Mul(psi_k340, Dot(&M, delta_mx) )
	sy := Mul(psi_k340, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek341() []float64 {

	sx := Mul(psi_k341, Dot(&M, delta_mx) )
	sy := Mul(psi_k341, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek342() []float64 {

	sx := Mul(psi_k342, Dot(&M, delta_mx) )
	sy := Mul(psi_k342, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek343() []float64 {

	sx := Mul(psi_k343, Dot(&M, delta_mx) )
	sy := Mul(psi_k343, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek344() []float64 {

	sx := Mul(psi_k344, Dot(&M, delta_mx) )
	sy := Mul(psi_k344, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek345() []float64 {

	sx := Mul(psi_k345, Dot(&M, delta_mx) )
	sy := Mul(psi_k345, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek346() []float64 {

	sx := Mul(psi_k346, Dot(&M, delta_mx) )
	sy := Mul(psi_k346, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek347() []float64 {

	sx := Mul(psi_k347, Dot(&M, delta_mx) )
	sy := Mul(psi_k347, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek348() []float64 {

	sx := Mul(psi_k348, Dot(&M, delta_mx) )
	sy := Mul(psi_k348, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek349() []float64 {

	sx := Mul(psi_k349, Dot(&M, delta_mx) )
	sy := Mul(psi_k349, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek350() []float64 {

	sx := Mul(psi_k350, Dot(&M, delta_mx) )
	sy := Mul(psi_k350, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek351() []float64 {

	sx := Mul(psi_k351, Dot(&M, delta_mx) )
	sy := Mul(psi_k351, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek352() []float64 {

	sx := Mul(psi_k352, Dot(&M, delta_mx) )
	sy := Mul(psi_k352, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek353() []float64 {

	sx := Mul(psi_k353, Dot(&M, delta_mx) )
	sy := Mul(psi_k353, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek354() []float64 {

	sx := Mul(psi_k354, Dot(&M, delta_mx) )
	sy := Mul(psi_k354, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek355() []float64 {

	sx := Mul(psi_k355, Dot(&M, delta_mx) )
	sy := Mul(psi_k355, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek356() []float64 {

	sx := Mul(psi_k356, Dot(&M, delta_mx) )
	sy := Mul(psi_k356, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek357() []float64 {

	sx := Mul(psi_k357, Dot(&M, delta_mx) )
	sy := Mul(psi_k357, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek358() []float64 {

	sx := Mul(psi_k358, Dot(&M, delta_mx) )
	sy := Mul(psi_k358, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek359() []float64 {

	sx := Mul(psi_k359, Dot(&M, delta_mx) )
	sy := Mul(psi_k359, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek360() []float64 {

	sx := Mul(psi_k360, Dot(&M, delta_mx) )
	sy := Mul(psi_k360, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek361() []float64 {

	sx := Mul(psi_k361, Dot(&M, delta_mx) )
	sy := Mul(psi_k361, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek362() []float64 {

	sx := Mul(psi_k362, Dot(&M, delta_mx) )
	sy := Mul(psi_k362, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek363() []float64 {

	sx := Mul(psi_k363, Dot(&M, delta_mx) )
	sy := Mul(psi_k363, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek364() []float64 {

	sx := Mul(psi_k364, Dot(&M, delta_mx) )
	sy := Mul(psi_k364, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek365() []float64 {

	sx := Mul(psi_k365, Dot(&M, delta_mx) )
	sy := Mul(psi_k365, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek366() []float64 {

	sx := Mul(psi_k366, Dot(&M, delta_mx) )
	sy := Mul(psi_k366, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek367() []float64 {

	sx := Mul(psi_k367, Dot(&M, delta_mx) )
	sy := Mul(psi_k367, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek368() []float64 {

	sx := Mul(psi_k368, Dot(&M, delta_mx) )
	sy := Mul(psi_k368, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek369() []float64 {

	sx := Mul(psi_k369, Dot(&M, delta_mx) )
	sy := Mul(psi_k369, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek370() []float64 {

	sx := Mul(psi_k370, Dot(&M, delta_mx) )
	sy := Mul(psi_k370, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek371() []float64 {

	sx := Mul(psi_k371, Dot(&M, delta_mx) )
	sy := Mul(psi_k371, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek372() []float64 {

	sx := Mul(psi_k372, Dot(&M, delta_mx) )
	sy := Mul(psi_k372, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek373() []float64 {

	sx := Mul(psi_k373, Dot(&M, delta_mx) )
	sy := Mul(psi_k373, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek374() []float64 {

	sx := Mul(psi_k374, Dot(&M, delta_mx) )
	sy := Mul(psi_k374, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek375() []float64 {

	sx := Mul(psi_k375, Dot(&M, delta_mx) )
	sy := Mul(psi_k375, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek376() []float64 {

	sx := Mul(psi_k376, Dot(&M, delta_mx) )
	sy := Mul(psi_k376, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek377() []float64 {

	sx := Mul(psi_k377, Dot(&M, delta_mx) )
	sy := Mul(psi_k377, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek378() []float64 {

	sx := Mul(psi_k378, Dot(&M, delta_mx) )
	sy := Mul(psi_k378, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek379() []float64 {

	sx := Mul(psi_k379, Dot(&M, delta_mx) )
	sy := Mul(psi_k379, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek380() []float64 {

	sx := Mul(psi_k380, Dot(&M, delta_mx) )
	sy := Mul(psi_k380, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek381() []float64 {

	sx := Mul(psi_k381, Dot(&M, delta_mx) )
	sy := Mul(psi_k381, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek382() []float64 {

	sx := Mul(psi_k382, Dot(&M, delta_mx) )
	sy := Mul(psi_k382, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek383() []float64 {

	sx := Mul(psi_k383, Dot(&M, delta_mx) )
	sy := Mul(psi_k383, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek384() []float64 {

	sx := Mul(psi_k384, Dot(&M, delta_mx) )
	sy := Mul(psi_k384, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek385() []float64 {

	sx := Mul(psi_k385, Dot(&M, delta_mx) )
	sy := Mul(psi_k385, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek386() []float64 {

	sx := Mul(psi_k386, Dot(&M, delta_mx) )
	sy := Mul(psi_k386, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek387() []float64 {

	sx := Mul(psi_k387, Dot(&M, delta_mx) )
	sy := Mul(psi_k387, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek388() []float64 {

	sx := Mul(psi_k388, Dot(&M, delta_mx) )
	sy := Mul(psi_k388, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek389() []float64 {

	sx := Mul(psi_k389, Dot(&M, delta_mx) )
	sy := Mul(psi_k389, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek390() []float64 {

	sx := Mul(psi_k390, Dot(&M, delta_mx) )
	sy := Mul(psi_k390, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek391() []float64 {

	sx := Mul(psi_k391, Dot(&M, delta_mx) )
	sy := Mul(psi_k391, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek392() []float64 {

	sx := Mul(psi_k392, Dot(&M, delta_mx) )
	sy := Mul(psi_k392, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek393() []float64 {

	sx := Mul(psi_k393, Dot(&M, delta_mx) )
	sy := Mul(psi_k393, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek394() []float64 {

	sx := Mul(psi_k394, Dot(&M, delta_mx) )
	sy := Mul(psi_k394, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek395() []float64 {

	sx := Mul(psi_k395, Dot(&M, delta_mx) )
	sy := Mul(psi_k395, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek396() []float64 {

	sx := Mul(psi_k396, Dot(&M, delta_mx) )
	sy := Mul(psi_k396, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek397() []float64 {

	sx := Mul(psi_k397, Dot(&M, delta_mx) )
	sy := Mul(psi_k397, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek398() []float64 {

	sx := Mul(psi_k398, Dot(&M, delta_mx) )
	sy := Mul(psi_k398, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek399() []float64 {

	sx := Mul(psi_k399, Dot(&M, delta_mx) )
	sy := Mul(psi_k399, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek400() []float64 {

	sx := Mul(psi_k400, Dot(&M, delta_mx) )
	sy := Mul(psi_k400, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek401() []float64 {

	sx := Mul(psi_k401, Dot(&M, delta_mx) )
	sy := Mul(psi_k401, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek402() []float64 {

	sx := Mul(psi_k402, Dot(&M, delta_mx) )
	sy := Mul(psi_k402, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek403() []float64 {

	sx := Mul(psi_k403, Dot(&M, delta_mx) )
	sy := Mul(psi_k403, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek404() []float64 {

	sx := Mul(psi_k404, Dot(&M, delta_mx) )
	sy := Mul(psi_k404, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek405() []float64 {

	sx := Mul(psi_k405, Dot(&M, delta_mx) )
	sy := Mul(psi_k405, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek406() []float64 {

	sx := Mul(psi_k406, Dot(&M, delta_mx) )
	sy := Mul(psi_k406, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek407() []float64 {

	sx := Mul(psi_k407, Dot(&M, delta_mx) )
	sy := Mul(psi_k407, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek408() []float64 {

	sx := Mul(psi_k408, Dot(&M, delta_mx) )
	sy := Mul(psi_k408, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek409() []float64 {

	sx := Mul(psi_k409, Dot(&M, delta_mx) )
	sy := Mul(psi_k409, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek410() []float64 {

	sx := Mul(psi_k410, Dot(&M, delta_mx) )
	sy := Mul(psi_k410, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek411() []float64 {

	sx := Mul(psi_k411, Dot(&M, delta_mx) )
	sy := Mul(psi_k411, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek412() []float64 {

	sx := Mul(psi_k412, Dot(&M, delta_mx) )
	sy := Mul(psi_k412, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek413() []float64 {

	sx := Mul(psi_k413, Dot(&M, delta_mx) )
	sy := Mul(psi_k413, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek414() []float64 {

	sx := Mul(psi_k414, Dot(&M, delta_mx) )
	sy := Mul(psi_k414, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek415() []float64 {

	sx := Mul(psi_k415, Dot(&M, delta_mx) )
	sy := Mul(psi_k415, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek416() []float64 {

	sx := Mul(psi_k416, Dot(&M, delta_mx) )
	sy := Mul(psi_k416, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek417() []float64 {

	sx := Mul(psi_k417, Dot(&M, delta_mx) )
	sy := Mul(psi_k417, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek418() []float64 {

	sx := Mul(psi_k418, Dot(&M, delta_mx) )
	sy := Mul(psi_k418, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek419() []float64 {

	sx := Mul(psi_k419, Dot(&M, delta_mx) )
	sy := Mul(psi_k419, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek420() []float64 {

	sx := Mul(psi_k420, Dot(&M, delta_mx) )
	sy := Mul(psi_k420, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek421() []float64 {

	sx := Mul(psi_k421, Dot(&M, delta_mx) )
	sy := Mul(psi_k421, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek422() []float64 {

	sx := Mul(psi_k422, Dot(&M, delta_mx) )
	sy := Mul(psi_k422, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek423() []float64 {

	sx := Mul(psi_k423, Dot(&M, delta_mx) )
	sy := Mul(psi_k423, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek424() []float64 {

	sx := Mul(psi_k424, Dot(&M, delta_mx) )
	sy := Mul(psi_k424, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek425() []float64 {

	sx := Mul(psi_k425, Dot(&M, delta_mx) )
	sy := Mul(psi_k425, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek426() []float64 {

	sx := Mul(psi_k426, Dot(&M, delta_mx) )
	sy := Mul(psi_k426, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek427() []float64 {

	sx := Mul(psi_k427, Dot(&M, delta_mx) )
	sy := Mul(psi_k427, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek428() []float64 {

	sx := Mul(psi_k428, Dot(&M, delta_mx) )
	sy := Mul(psi_k428, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek429() []float64 {

	sx := Mul(psi_k429, Dot(&M, delta_mx) )
	sy := Mul(psi_k429, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek430() []float64 {

	sx := Mul(psi_k430, Dot(&M, delta_mx) )
	sy := Mul(psi_k430, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek431() []float64 {

	sx := Mul(psi_k431, Dot(&M, delta_mx) )
	sy := Mul(psi_k431, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek432() []float64 {

	sx := Mul(psi_k432, Dot(&M, delta_mx) )
	sy := Mul(psi_k432, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek433() []float64 {

	sx := Mul(psi_k433, Dot(&M, delta_mx) )
	sy := Mul(psi_k433, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek434() []float64 {

	sx := Mul(psi_k434, Dot(&M, delta_mx) )
	sy := Mul(psi_k434, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek435() []float64 {

	sx := Mul(psi_k435, Dot(&M, delta_mx) )
	sy := Mul(psi_k435, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek436() []float64 {

	sx := Mul(psi_k436, Dot(&M, delta_mx) )
	sy := Mul(psi_k436, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek437() []float64 {

	sx := Mul(psi_k437, Dot(&M, delta_mx) )
	sy := Mul(psi_k437, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek438() []float64 {

	sx := Mul(psi_k438, Dot(&M, delta_mx) )
	sy := Mul(psi_k438, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek439() []float64 {

	sx := Mul(psi_k439, Dot(&M, delta_mx) )
	sy := Mul(psi_k439, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek440() []float64 {

	sx := Mul(psi_k440, Dot(&M, delta_mx) )
	sy := Mul(psi_k440, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek441() []float64 {

	sx := Mul(psi_k441, Dot(&M, delta_mx) )
	sy := Mul(psi_k441, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek442() []float64 {

	sx := Mul(psi_k442, Dot(&M, delta_mx) )
	sy := Mul(psi_k442, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek443() []float64 {

	sx := Mul(psi_k443, Dot(&M, delta_mx) )
	sy := Mul(psi_k443, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek444() []float64 {

	sx := Mul(psi_k444, Dot(&M, delta_mx) )
	sy := Mul(psi_k444, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek445() []float64 {

	sx := Mul(psi_k445, Dot(&M, delta_mx) )
	sy := Mul(psi_k445, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek446() []float64 {

	sx := Mul(psi_k446, Dot(&M, delta_mx) )
	sy := Mul(psi_k446, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek447() []float64 {

	sx := Mul(psi_k447, Dot(&M, delta_mx) )
	sy := Mul(psi_k447, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek448() []float64 {

	sx := Mul(psi_k448, Dot(&M, delta_mx) )
	sy := Mul(psi_k448, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek449() []float64 {

	sx := Mul(psi_k449, Dot(&M, delta_mx) )
	sy := Mul(psi_k449, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek450() []float64 {

	sx := Mul(psi_k450, Dot(&M, delta_mx) )
	sy := Mul(psi_k450, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek451() []float64 {

	sx := Mul(psi_k451, Dot(&M, delta_mx) )
	sy := Mul(psi_k451, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek452() []float64 {

	sx := Mul(psi_k452, Dot(&M, delta_mx) )
	sy := Mul(psi_k452, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek453() []float64 {

	sx := Mul(psi_k453, Dot(&M, delta_mx) )
	sy := Mul(psi_k453, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek454() []float64 {

	sx := Mul(psi_k454, Dot(&M, delta_mx) )
	sy := Mul(psi_k454, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek455() []float64 {

	sx := Mul(psi_k455, Dot(&M, delta_mx) )
	sy := Mul(psi_k455, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek456() []float64 {

	sx := Mul(psi_k456, Dot(&M, delta_mx) )
	sy := Mul(psi_k456, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek457() []float64 {

	sx := Mul(psi_k457, Dot(&M, delta_mx) )
	sy := Mul(psi_k457, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek458() []float64 {

	sx := Mul(psi_k458, Dot(&M, delta_mx) )
	sy := Mul(psi_k458, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek459() []float64 {

	sx := Mul(psi_k459, Dot(&M, delta_mx) )
	sy := Mul(psi_k459, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek460() []float64 {

	sx := Mul(psi_k460, Dot(&M, delta_mx) )
	sy := Mul(psi_k460, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek461() []float64 {

	sx := Mul(psi_k461, Dot(&M, delta_mx) )
	sy := Mul(psi_k461, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek462() []float64 {

	sx := Mul(psi_k462, Dot(&M, delta_mx) )
	sy := Mul(psi_k462, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek463() []float64 {

	sx := Mul(psi_k463, Dot(&M, delta_mx) )
	sy := Mul(psi_k463, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek464() []float64 {

	sx := Mul(psi_k464, Dot(&M, delta_mx) )
	sy := Mul(psi_k464, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek465() []float64 {

	sx := Mul(psi_k465, Dot(&M, delta_mx) )
	sy := Mul(psi_k465, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek466() []float64 {

	sx := Mul(psi_k466, Dot(&M, delta_mx) )
	sy := Mul(psi_k466, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek467() []float64 {

	sx := Mul(psi_k467, Dot(&M, delta_mx) )
	sy := Mul(psi_k467, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek468() []float64 {

	sx := Mul(psi_k468, Dot(&M, delta_mx) )
	sy := Mul(psi_k468, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek469() []float64 {

	sx := Mul(psi_k469, Dot(&M, delta_mx) )
	sy := Mul(psi_k469, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek470() []float64 {

	sx := Mul(psi_k470, Dot(&M, delta_mx) )
	sy := Mul(psi_k470, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek471() []float64 {

	sx := Mul(psi_k471, Dot(&M, delta_mx) )
	sy := Mul(psi_k471, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek472() []float64 {

	sx := Mul(psi_k472, Dot(&M, delta_mx) )
	sy := Mul(psi_k472, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek473() []float64 {

	sx := Mul(psi_k473, Dot(&M, delta_mx) )
	sy := Mul(psi_k473, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek474() []float64 {

	sx := Mul(psi_k474, Dot(&M, delta_mx) )
	sy := Mul(psi_k474, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek475() []float64 {

	sx := Mul(psi_k475, Dot(&M, delta_mx) )
	sy := Mul(psi_k475, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek476() []float64 {

	sx := Mul(psi_k476, Dot(&M, delta_mx) )
	sy := Mul(psi_k476, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek477() []float64 {

	sx := Mul(psi_k477, Dot(&M, delta_mx) )
	sy := Mul(psi_k477, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek478() []float64 {

	sx := Mul(psi_k478, Dot(&M, delta_mx) )
	sy := Mul(psi_k478, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek479() []float64 {

	sx := Mul(psi_k479, Dot(&M, delta_mx) )
	sy := Mul(psi_k479, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek480() []float64 {

	sx := Mul(psi_k480, Dot(&M, delta_mx) )
	sy := Mul(psi_k480, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek481() []float64 {

	sx := Mul(psi_k481, Dot(&M, delta_mx) )
	sy := Mul(psi_k481, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek482() []float64 {

	sx := Mul(psi_k482, Dot(&M, delta_mx) )
	sy := Mul(psi_k482, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek483() []float64 {

	sx := Mul(psi_k483, Dot(&M, delta_mx) )
	sy := Mul(psi_k483, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek484() []float64 {

	sx := Mul(psi_k484, Dot(&M, delta_mx) )
	sy := Mul(psi_k484, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek485() []float64 {

	sx := Mul(psi_k485, Dot(&M, delta_mx) )
	sy := Mul(psi_k485, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek486() []float64 {

	sx := Mul(psi_k486, Dot(&M, delta_mx) )
	sy := Mul(psi_k486, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek487() []float64 {

	sx := Mul(psi_k487, Dot(&M, delta_mx) )
	sy := Mul(psi_k487, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek488() []float64 {

	sx := Mul(psi_k488, Dot(&M, delta_mx) )
	sy := Mul(psi_k488, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek489() []float64 {

	sx := Mul(psi_k489, Dot(&M, delta_mx) )
	sy := Mul(psi_k489, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek490() []float64 {

	sx := Mul(psi_k490, Dot(&M, delta_mx) )
	sy := Mul(psi_k490, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek491() []float64 {

	sx := Mul(psi_k491, Dot(&M, delta_mx) )
	sy := Mul(psi_k491, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek492() []float64 {

	sx := Mul(psi_k492, Dot(&M, delta_mx) )
	sy := Mul(psi_k492, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek493() []float64 {

	sx := Mul(psi_k493, Dot(&M, delta_mx) )
	sy := Mul(psi_k493, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek494() []float64 {

	sx := Mul(psi_k494, Dot(&M, delta_mx) )
	sy := Mul(psi_k494, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek495() []float64 {

	sx := Mul(psi_k495, Dot(&M, delta_mx) )
	sy := Mul(psi_k495, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek496() []float64 {

	sx := Mul(psi_k496, Dot(&M, delta_mx) )
	sy := Mul(psi_k496, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek497() []float64 {

	sx := Mul(psi_k497, Dot(&M, delta_mx) )
	sy := Mul(psi_k497, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek498() []float64 {

	sx := Mul(psi_k498, Dot(&M, delta_mx) )
	sy := Mul(psi_k498, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}

func GetModeAmplitudek499() []float64 {

	sx := Mul(psi_k499, Dot(&M, delta_mx) )
	sy := Mul(psi_k499, Dot(&M, delta_my) )
		
	wx := ValueOf(sx)
	defer cuda.Recycle(wx)	

	wy := ValueOf(sy)
	defer cuda.Recycle(wy)	

	amp := make([]float64, 3)

	amp[0] = float64( cuda.Sum(wx) )
	amp[1] = float64( cuda.Sum(wy) )
	amp[2] = float64( 0.0 )

	return amp
}