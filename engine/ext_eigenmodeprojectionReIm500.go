package engine

// **************************************** 
// Author(s): Joo-Von Kim, C2N, CNRS/Univ. Paris-Saclay
//
// This module projects the magnetization onto user-supplied eigenvectors
// Re(psi_k) and Im(psi_k), which are both 3-component fields. The module
// returns the amplitudes
// 
// 	Re(b_k) = int_dV {(Im(psi_k) x m0).(m-m0)}
//	Im(b_k) = int_dV {(Re(psi_k) x m0).(m-m0)}
//  n_k     = b_k^* b_k = Re(b_k)^2 + Im(b_k)^2
//
// The user-supplied vector fields can be added in the source .mx3 file with
//  M0.Add( LoadFile(("m0_file.ovf"),1) )	
//	psiRe_k.Add( LoadFile(("psi_file.ovf"),1) )	
//	psiIm_k.Add( LoadFile(("psi_file.ovf"),1) )	
//	etc.
//
// Acknowledgements:
// This work was supported by Horizon 2020 Research Framework Programme of the
// European Commission under grant agreement No. 899646 (k-Net). 
//
// **************************************** 


import (
	"github.com/mumax/3/cuda"
)


var (
	M0				= NewExcitation("M0", "", "Equilibrium magnetization configuration")

	psiRe_k000		= NewExcitation("psiRe_k000", "", "Eigenmode spatial profile")
	psiRe_k001		= NewExcitation("psiRe_k001", "", "Eigenmode spatial profile")
	psiRe_k002		= NewExcitation("psiRe_k002", "", "Eigenmode spatial profile")
	psiRe_k003		= NewExcitation("psiRe_k003", "", "Eigenmode spatial profile")
	psiRe_k004		= NewExcitation("psiRe_k004", "", "Eigenmode spatial profile")
	psiRe_k005		= NewExcitation("psiRe_k005", "", "Eigenmode spatial profile")
	psiRe_k006		= NewExcitation("psiRe_k006", "", "Eigenmode spatial profile")
	psiRe_k007		= NewExcitation("psiRe_k007", "", "Eigenmode spatial profile")
	psiRe_k008		= NewExcitation("psiRe_k008", "", "Eigenmode spatial profile")
	psiRe_k009		= NewExcitation("psiRe_k009", "", "Eigenmode spatial profile")
	psiRe_k010		= NewExcitation("psiRe_k010", "", "Eigenmode spatial profile")
	psiRe_k011		= NewExcitation("psiRe_k011", "", "Eigenmode spatial profile")
	psiRe_k012		= NewExcitation("psiRe_k012", "", "Eigenmode spatial profile")
	psiRe_k013		= NewExcitation("psiRe_k013", "", "Eigenmode spatial profile")
	psiRe_k014		= NewExcitation("psiRe_k014", "", "Eigenmode spatial profile")
	psiRe_k015		= NewExcitation("psiRe_k015", "", "Eigenmode spatial profile")
	psiRe_k016		= NewExcitation("psiRe_k016", "", "Eigenmode spatial profile")
	psiRe_k017		= NewExcitation("psiRe_k017", "", "Eigenmode spatial profile")
	psiRe_k018		= NewExcitation("psiRe_k018", "", "Eigenmode spatial profile")
	psiRe_k019		= NewExcitation("psiRe_k019", "", "Eigenmode spatial profile")
	psiRe_k020		= NewExcitation("psiRe_k020", "", "Eigenmode spatial profile")
	psiRe_k021		= NewExcitation("psiRe_k021", "", "Eigenmode spatial profile")
	psiRe_k022		= NewExcitation("psiRe_k022", "", "Eigenmode spatial profile")
	psiRe_k023		= NewExcitation("psiRe_k023", "", "Eigenmode spatial profile")
	psiRe_k024		= NewExcitation("psiRe_k024", "", "Eigenmode spatial profile")
	psiRe_k025		= NewExcitation("psiRe_k025", "", "Eigenmode spatial profile")
	psiRe_k026		= NewExcitation("psiRe_k026", "", "Eigenmode spatial profile")
	psiRe_k027		= NewExcitation("psiRe_k027", "", "Eigenmode spatial profile")
	psiRe_k028		= NewExcitation("psiRe_k028", "", "Eigenmode spatial profile")
	psiRe_k029		= NewExcitation("psiRe_k029", "", "Eigenmode spatial profile")
	psiRe_k030		= NewExcitation("psiRe_k030", "", "Eigenmode spatial profile")
	psiRe_k031		= NewExcitation("psiRe_k031", "", "Eigenmode spatial profile")
	psiRe_k032		= NewExcitation("psiRe_k032", "", "Eigenmode spatial profile")
	psiRe_k033		= NewExcitation("psiRe_k033", "", "Eigenmode spatial profile")
	psiRe_k034		= NewExcitation("psiRe_k034", "", "Eigenmode spatial profile")
	psiRe_k035		= NewExcitation("psiRe_k035", "", "Eigenmode spatial profile")
	psiRe_k036		= NewExcitation("psiRe_k036", "", "Eigenmode spatial profile")
	psiRe_k037		= NewExcitation("psiRe_k037", "", "Eigenmode spatial profile")
	psiRe_k038		= NewExcitation("psiRe_k038", "", "Eigenmode spatial profile")
	psiRe_k039		= NewExcitation("psiRe_k039", "", "Eigenmode spatial profile")
	psiRe_k040		= NewExcitation("psiRe_k040", "", "Eigenmode spatial profile")
	psiRe_k041		= NewExcitation("psiRe_k041", "", "Eigenmode spatial profile")
	psiRe_k042		= NewExcitation("psiRe_k042", "", "Eigenmode spatial profile")
	psiRe_k043		= NewExcitation("psiRe_k043", "", "Eigenmode spatial profile")
	psiRe_k044		= NewExcitation("psiRe_k044", "", "Eigenmode spatial profile")
	psiRe_k045		= NewExcitation("psiRe_k045", "", "Eigenmode spatial profile")
	psiRe_k046		= NewExcitation("psiRe_k046", "", "Eigenmode spatial profile")
	psiRe_k047		= NewExcitation("psiRe_k047", "", "Eigenmode spatial profile")
	psiRe_k048		= NewExcitation("psiRe_k048", "", "Eigenmode spatial profile")
	psiRe_k049		= NewExcitation("psiRe_k049", "", "Eigenmode spatial profile")
	psiRe_k050		= NewExcitation("psiRe_k050", "", "Eigenmode spatial profile")
	psiRe_k051		= NewExcitation("psiRe_k051", "", "Eigenmode spatial profile")
	psiRe_k052		= NewExcitation("psiRe_k052", "", "Eigenmode spatial profile")
	psiRe_k053		= NewExcitation("psiRe_k053", "", "Eigenmode spatial profile")
	psiRe_k054		= NewExcitation("psiRe_k054", "", "Eigenmode spatial profile")
	psiRe_k055		= NewExcitation("psiRe_k055", "", "Eigenmode spatial profile")
	psiRe_k056		= NewExcitation("psiRe_k056", "", "Eigenmode spatial profile")
	psiRe_k057		= NewExcitation("psiRe_k057", "", "Eigenmode spatial profile")
	psiRe_k058		= NewExcitation("psiRe_k058", "", "Eigenmode spatial profile")
	psiRe_k059		= NewExcitation("psiRe_k059", "", "Eigenmode spatial profile")
	psiRe_k060		= NewExcitation("psiRe_k060", "", "Eigenmode spatial profile")
	psiRe_k061		= NewExcitation("psiRe_k061", "", "Eigenmode spatial profile")
	psiRe_k062		= NewExcitation("psiRe_k062", "", "Eigenmode spatial profile")
	psiRe_k063		= NewExcitation("psiRe_k063", "", "Eigenmode spatial profile")
	psiRe_k064		= NewExcitation("psiRe_k064", "", "Eigenmode spatial profile")
	psiRe_k065		= NewExcitation("psiRe_k065", "", "Eigenmode spatial profile")
	psiRe_k066		= NewExcitation("psiRe_k066", "", "Eigenmode spatial profile")
	psiRe_k067		= NewExcitation("psiRe_k067", "", "Eigenmode spatial profile")
	psiRe_k068		= NewExcitation("psiRe_k068", "", "Eigenmode spatial profile")
	psiRe_k069		= NewExcitation("psiRe_k069", "", "Eigenmode spatial profile")
	psiRe_k070		= NewExcitation("psiRe_k070", "", "Eigenmode spatial profile")
	psiRe_k071		= NewExcitation("psiRe_k071", "", "Eigenmode spatial profile")
	psiRe_k072		= NewExcitation("psiRe_k072", "", "Eigenmode spatial profile")
	psiRe_k073		= NewExcitation("psiRe_k073", "", "Eigenmode spatial profile")
	psiRe_k074		= NewExcitation("psiRe_k074", "", "Eigenmode spatial profile")
	psiRe_k075		= NewExcitation("psiRe_k075", "", "Eigenmode spatial profile")
	psiRe_k076		= NewExcitation("psiRe_k076", "", "Eigenmode spatial profile")
	psiRe_k077		= NewExcitation("psiRe_k077", "", "Eigenmode spatial profile")
	psiRe_k078		= NewExcitation("psiRe_k078", "", "Eigenmode spatial profile")
	psiRe_k079		= NewExcitation("psiRe_k079", "", "Eigenmode spatial profile")
	psiRe_k080		= NewExcitation("psiRe_k080", "", "Eigenmode spatial profile")
	psiRe_k081		= NewExcitation("psiRe_k081", "", "Eigenmode spatial profile")
	psiRe_k082		= NewExcitation("psiRe_k082", "", "Eigenmode spatial profile")
	psiRe_k083		= NewExcitation("psiRe_k083", "", "Eigenmode spatial profile")
	psiRe_k084		= NewExcitation("psiRe_k084", "", "Eigenmode spatial profile")
	psiRe_k085		= NewExcitation("psiRe_k085", "", "Eigenmode spatial profile")
	psiRe_k086		= NewExcitation("psiRe_k086", "", "Eigenmode spatial profile")
	psiRe_k087		= NewExcitation("psiRe_k087", "", "Eigenmode spatial profile")
	psiRe_k088		= NewExcitation("psiRe_k088", "", "Eigenmode spatial profile")
	psiRe_k089		= NewExcitation("psiRe_k089", "", "Eigenmode spatial profile")
	psiRe_k090		= NewExcitation("psiRe_k090", "", "Eigenmode spatial profile")
	psiRe_k091		= NewExcitation("psiRe_k091", "", "Eigenmode spatial profile")
	psiRe_k092		= NewExcitation("psiRe_k092", "", "Eigenmode spatial profile")
	psiRe_k093		= NewExcitation("psiRe_k093", "", "Eigenmode spatial profile")
	psiRe_k094		= NewExcitation("psiRe_k094", "", "Eigenmode spatial profile")
	psiRe_k095		= NewExcitation("psiRe_k095", "", "Eigenmode spatial profile")
	psiRe_k096		= NewExcitation("psiRe_k096", "", "Eigenmode spatial profile")
	psiRe_k097		= NewExcitation("psiRe_k097", "", "Eigenmode spatial profile")
	psiRe_k098		= NewExcitation("psiRe_k098", "", "Eigenmode spatial profile")
	psiRe_k099		= NewExcitation("psiRe_k099", "", "Eigenmode spatial profile")
	psiRe_k100		= NewExcitation("psiRe_k100", "", "Eigenmode spatial profile")
	psiRe_k101		= NewExcitation("psiRe_k101", "", "Eigenmode spatial profile")
	psiRe_k102		= NewExcitation("psiRe_k102", "", "Eigenmode spatial profile")
	psiRe_k103		= NewExcitation("psiRe_k103", "", "Eigenmode spatial profile")
	psiRe_k104		= NewExcitation("psiRe_k104", "", "Eigenmode spatial profile")
	psiRe_k105		= NewExcitation("psiRe_k105", "", "Eigenmode spatial profile")
	psiRe_k106		= NewExcitation("psiRe_k106", "", "Eigenmode spatial profile")
	psiRe_k107		= NewExcitation("psiRe_k107", "", "Eigenmode spatial profile")
	psiRe_k108		= NewExcitation("psiRe_k108", "", "Eigenmode spatial profile")
	psiRe_k109		= NewExcitation("psiRe_k109", "", "Eigenmode spatial profile")
	psiRe_k110		= NewExcitation("psiRe_k110", "", "Eigenmode spatial profile")
	psiRe_k111		= NewExcitation("psiRe_k111", "", "Eigenmode spatial profile")
	psiRe_k112		= NewExcitation("psiRe_k112", "", "Eigenmode spatial profile")
	psiRe_k113		= NewExcitation("psiRe_k113", "", "Eigenmode spatial profile")
	psiRe_k114		= NewExcitation("psiRe_k114", "", "Eigenmode spatial profile")
	psiRe_k115		= NewExcitation("psiRe_k115", "", "Eigenmode spatial profile")
	psiRe_k116		= NewExcitation("psiRe_k116", "", "Eigenmode spatial profile")
	psiRe_k117		= NewExcitation("psiRe_k117", "", "Eigenmode spatial profile")
	psiRe_k118		= NewExcitation("psiRe_k118", "", "Eigenmode spatial profile")
	psiRe_k119		= NewExcitation("psiRe_k119", "", "Eigenmode spatial profile")
	psiRe_k120		= NewExcitation("psiRe_k120", "", "Eigenmode spatial profile")
	psiRe_k121		= NewExcitation("psiRe_k121", "", "Eigenmode spatial profile")
	psiRe_k122		= NewExcitation("psiRe_k122", "", "Eigenmode spatial profile")
	psiRe_k123		= NewExcitation("psiRe_k123", "", "Eigenmode spatial profile")
	psiRe_k124		= NewExcitation("psiRe_k124", "", "Eigenmode spatial profile")
	psiRe_k125		= NewExcitation("psiRe_k125", "", "Eigenmode spatial profile")
	psiRe_k126		= NewExcitation("psiRe_k126", "", "Eigenmode spatial profile")
	psiRe_k127		= NewExcitation("psiRe_k127", "", "Eigenmode spatial profile")
	psiRe_k128		= NewExcitation("psiRe_k128", "", "Eigenmode spatial profile")
	psiRe_k129		= NewExcitation("psiRe_k129", "", "Eigenmode spatial profile")
	psiRe_k130		= NewExcitation("psiRe_k130", "", "Eigenmode spatial profile")
	psiRe_k131		= NewExcitation("psiRe_k131", "", "Eigenmode spatial profile")
	psiRe_k132		= NewExcitation("psiRe_k132", "", "Eigenmode spatial profile")
	psiRe_k133		= NewExcitation("psiRe_k133", "", "Eigenmode spatial profile")
	psiRe_k134		= NewExcitation("psiRe_k134", "", "Eigenmode spatial profile")
	psiRe_k135		= NewExcitation("psiRe_k135", "", "Eigenmode spatial profile")
	psiRe_k136		= NewExcitation("psiRe_k136", "", "Eigenmode spatial profile")
	psiRe_k137		= NewExcitation("psiRe_k137", "", "Eigenmode spatial profile")
	psiRe_k138		= NewExcitation("psiRe_k138", "", "Eigenmode spatial profile")
	psiRe_k139		= NewExcitation("psiRe_k139", "", "Eigenmode spatial profile")
	psiRe_k140		= NewExcitation("psiRe_k140", "", "Eigenmode spatial profile")
	psiRe_k141		= NewExcitation("psiRe_k141", "", "Eigenmode spatial profile")
	psiRe_k142		= NewExcitation("psiRe_k142", "", "Eigenmode spatial profile")
	psiRe_k143		= NewExcitation("psiRe_k143", "", "Eigenmode spatial profile")
	psiRe_k144		= NewExcitation("psiRe_k144", "", "Eigenmode spatial profile")
	psiRe_k145		= NewExcitation("psiRe_k145", "", "Eigenmode spatial profile")
	psiRe_k146		= NewExcitation("psiRe_k146", "", "Eigenmode spatial profile")
	psiRe_k147		= NewExcitation("psiRe_k147", "", "Eigenmode spatial profile")
	psiRe_k148		= NewExcitation("psiRe_k148", "", "Eigenmode spatial profile")
	psiRe_k149		= NewExcitation("psiRe_k149", "", "Eigenmode spatial profile")
	psiRe_k150		= NewExcitation("psiRe_k150", "", "Eigenmode spatial profile")
	psiRe_k151		= NewExcitation("psiRe_k151", "", "Eigenmode spatial profile")
	psiRe_k152		= NewExcitation("psiRe_k152", "", "Eigenmode spatial profile")
	psiRe_k153		= NewExcitation("psiRe_k153", "", "Eigenmode spatial profile")
	psiRe_k154		= NewExcitation("psiRe_k154", "", "Eigenmode spatial profile")
	psiRe_k155		= NewExcitation("psiRe_k155", "", "Eigenmode spatial profile")
	psiRe_k156		= NewExcitation("psiRe_k156", "", "Eigenmode spatial profile")
	psiRe_k157		= NewExcitation("psiRe_k157", "", "Eigenmode spatial profile")
	psiRe_k158		= NewExcitation("psiRe_k158", "", "Eigenmode spatial profile")
	psiRe_k159		= NewExcitation("psiRe_k159", "", "Eigenmode spatial profile")
	psiRe_k160		= NewExcitation("psiRe_k160", "", "Eigenmode spatial profile")
	psiRe_k161		= NewExcitation("psiRe_k161", "", "Eigenmode spatial profile")
	psiRe_k162		= NewExcitation("psiRe_k162", "", "Eigenmode spatial profile")
	psiRe_k163		= NewExcitation("psiRe_k163", "", "Eigenmode spatial profile")
	psiRe_k164		= NewExcitation("psiRe_k164", "", "Eigenmode spatial profile")
	psiRe_k165		= NewExcitation("psiRe_k165", "", "Eigenmode spatial profile")
	psiRe_k166		= NewExcitation("psiRe_k166", "", "Eigenmode spatial profile")
	psiRe_k167		= NewExcitation("psiRe_k167", "", "Eigenmode spatial profile")
	psiRe_k168		= NewExcitation("psiRe_k168", "", "Eigenmode spatial profile")
	psiRe_k169		= NewExcitation("psiRe_k169", "", "Eigenmode spatial profile")
	psiRe_k170		= NewExcitation("psiRe_k170", "", "Eigenmode spatial profile")
	psiRe_k171		= NewExcitation("psiRe_k171", "", "Eigenmode spatial profile")
	psiRe_k172		= NewExcitation("psiRe_k172", "", "Eigenmode spatial profile")
	psiRe_k173		= NewExcitation("psiRe_k173", "", "Eigenmode spatial profile")
	psiRe_k174		= NewExcitation("psiRe_k174", "", "Eigenmode spatial profile")
	psiRe_k175		= NewExcitation("psiRe_k175", "", "Eigenmode spatial profile")
	psiRe_k176		= NewExcitation("psiRe_k176", "", "Eigenmode spatial profile")
	psiRe_k177		= NewExcitation("psiRe_k177", "", "Eigenmode spatial profile")
	psiRe_k178		= NewExcitation("psiRe_k178", "", "Eigenmode spatial profile")
	psiRe_k179		= NewExcitation("psiRe_k179", "", "Eigenmode spatial profile")
	psiRe_k180		= NewExcitation("psiRe_k180", "", "Eigenmode spatial profile")
	psiRe_k181		= NewExcitation("psiRe_k181", "", "Eigenmode spatial profile")
	psiRe_k182		= NewExcitation("psiRe_k182", "", "Eigenmode spatial profile")
	psiRe_k183		= NewExcitation("psiRe_k183", "", "Eigenmode spatial profile")
	psiRe_k184		= NewExcitation("psiRe_k184", "", "Eigenmode spatial profile")
	psiRe_k185		= NewExcitation("psiRe_k185", "", "Eigenmode spatial profile")
	psiRe_k186		= NewExcitation("psiRe_k186", "", "Eigenmode spatial profile")
	psiRe_k187		= NewExcitation("psiRe_k187", "", "Eigenmode spatial profile")
	psiRe_k188		= NewExcitation("psiRe_k188", "", "Eigenmode spatial profile")
	psiRe_k189		= NewExcitation("psiRe_k189", "", "Eigenmode spatial profile")
	psiRe_k190		= NewExcitation("psiRe_k190", "", "Eigenmode spatial profile")
	psiRe_k191		= NewExcitation("psiRe_k191", "", "Eigenmode spatial profile")
	psiRe_k192		= NewExcitation("psiRe_k192", "", "Eigenmode spatial profile")
	psiRe_k193		= NewExcitation("psiRe_k193", "", "Eigenmode spatial profile")
	psiRe_k194		= NewExcitation("psiRe_k194", "", "Eigenmode spatial profile")
	psiRe_k195		= NewExcitation("psiRe_k195", "", "Eigenmode spatial profile")
	psiRe_k196		= NewExcitation("psiRe_k196", "", "Eigenmode spatial profile")
	psiRe_k197		= NewExcitation("psiRe_k197", "", "Eigenmode spatial profile")
	psiRe_k198		= NewExcitation("psiRe_k198", "", "Eigenmode spatial profile")
	psiRe_k199		= NewExcitation("psiRe_k199", "", "Eigenmode spatial profile")
	psiRe_k200		= NewExcitation("psiRe_k200", "", "Eigenmode spatial profile")
	psiRe_k201		= NewExcitation("psiRe_k201", "", "Eigenmode spatial profile")
	psiRe_k202		= NewExcitation("psiRe_k202", "", "Eigenmode spatial profile")
	psiRe_k203		= NewExcitation("psiRe_k203", "", "Eigenmode spatial profile")
	psiRe_k204		= NewExcitation("psiRe_k204", "", "Eigenmode spatial profile")
	psiRe_k205		= NewExcitation("psiRe_k205", "", "Eigenmode spatial profile")
	psiRe_k206		= NewExcitation("psiRe_k206", "", "Eigenmode spatial profile")
	psiRe_k207		= NewExcitation("psiRe_k207", "", "Eigenmode spatial profile")
	psiRe_k208		= NewExcitation("psiRe_k208", "", "Eigenmode spatial profile")
	psiRe_k209		= NewExcitation("psiRe_k209", "", "Eigenmode spatial profile")
	psiRe_k210		= NewExcitation("psiRe_k210", "", "Eigenmode spatial profile")
	psiRe_k211		= NewExcitation("psiRe_k211", "", "Eigenmode spatial profile")
	psiRe_k212		= NewExcitation("psiRe_k212", "", "Eigenmode spatial profile")
	psiRe_k213		= NewExcitation("psiRe_k213", "", "Eigenmode spatial profile")
	psiRe_k214		= NewExcitation("psiRe_k214", "", "Eigenmode spatial profile")
	psiRe_k215		= NewExcitation("psiRe_k215", "", "Eigenmode spatial profile")
	psiRe_k216		= NewExcitation("psiRe_k216", "", "Eigenmode spatial profile")
	psiRe_k217		= NewExcitation("psiRe_k217", "", "Eigenmode spatial profile")
	psiRe_k218		= NewExcitation("psiRe_k218", "", "Eigenmode spatial profile")
	psiRe_k219		= NewExcitation("psiRe_k219", "", "Eigenmode spatial profile")
	psiRe_k220		= NewExcitation("psiRe_k220", "", "Eigenmode spatial profile")
	psiRe_k221		= NewExcitation("psiRe_k221", "", "Eigenmode spatial profile")
	psiRe_k222		= NewExcitation("psiRe_k222", "", "Eigenmode spatial profile")
	psiRe_k223		= NewExcitation("psiRe_k223", "", "Eigenmode spatial profile")
	psiRe_k224		= NewExcitation("psiRe_k224", "", "Eigenmode spatial profile")
	psiRe_k225		= NewExcitation("psiRe_k225", "", "Eigenmode spatial profile")
	psiRe_k226		= NewExcitation("psiRe_k226", "", "Eigenmode spatial profile")
	psiRe_k227		= NewExcitation("psiRe_k227", "", "Eigenmode spatial profile")
	psiRe_k228		= NewExcitation("psiRe_k228", "", "Eigenmode spatial profile")
	psiRe_k229		= NewExcitation("psiRe_k229", "", "Eigenmode spatial profile")
	psiRe_k230		= NewExcitation("psiRe_k230", "", "Eigenmode spatial profile")
	psiRe_k231		= NewExcitation("psiRe_k231", "", "Eigenmode spatial profile")
	psiRe_k232		= NewExcitation("psiRe_k232", "", "Eigenmode spatial profile")
	psiRe_k233		= NewExcitation("psiRe_k233", "", "Eigenmode spatial profile")
	psiRe_k234		= NewExcitation("psiRe_k234", "", "Eigenmode spatial profile")
	psiRe_k235		= NewExcitation("psiRe_k235", "", "Eigenmode spatial profile")
	psiRe_k236		= NewExcitation("psiRe_k236", "", "Eigenmode spatial profile")
	psiRe_k237		= NewExcitation("psiRe_k237", "", "Eigenmode spatial profile")
	psiRe_k238		= NewExcitation("psiRe_k238", "", "Eigenmode spatial profile")
	psiRe_k239		= NewExcitation("psiRe_k239", "", "Eigenmode spatial profile")
	psiRe_k240		= NewExcitation("psiRe_k240", "", "Eigenmode spatial profile")
	psiRe_k241		= NewExcitation("psiRe_k241", "", "Eigenmode spatial profile")
	psiRe_k242		= NewExcitation("psiRe_k242", "", "Eigenmode spatial profile")
	psiRe_k243		= NewExcitation("psiRe_k243", "", "Eigenmode spatial profile")
	psiRe_k244		= NewExcitation("psiRe_k244", "", "Eigenmode spatial profile")
	psiRe_k245		= NewExcitation("psiRe_k245", "", "Eigenmode spatial profile")
	psiRe_k246		= NewExcitation("psiRe_k246", "", "Eigenmode spatial profile")
	psiRe_k247		= NewExcitation("psiRe_k247", "", "Eigenmode spatial profile")
	psiRe_k248		= NewExcitation("psiRe_k248", "", "Eigenmode spatial profile")
	psiRe_k249		= NewExcitation("psiRe_k249", "", "Eigenmode spatial profile")
	psiRe_k250		= NewExcitation("psiRe_k250", "", "Eigenmode spatial profile")
	psiRe_k251		= NewExcitation("psiRe_k251", "", "Eigenmode spatial profile")
	psiRe_k252		= NewExcitation("psiRe_k252", "", "Eigenmode spatial profile")
	psiRe_k253		= NewExcitation("psiRe_k253", "", "Eigenmode spatial profile")
	psiRe_k254		= NewExcitation("psiRe_k254", "", "Eigenmode spatial profile")
	psiRe_k255		= NewExcitation("psiRe_k255", "", "Eigenmode spatial profile")
	psiRe_k256		= NewExcitation("psiRe_k256", "", "Eigenmode spatial profile")
	psiRe_k257		= NewExcitation("psiRe_k257", "", "Eigenmode spatial profile")
	psiRe_k258		= NewExcitation("psiRe_k258", "", "Eigenmode spatial profile")
	psiRe_k259		= NewExcitation("psiRe_k259", "", "Eigenmode spatial profile")
	psiRe_k260		= NewExcitation("psiRe_k260", "", "Eigenmode spatial profile")
	psiRe_k261		= NewExcitation("psiRe_k261", "", "Eigenmode spatial profile")
	psiRe_k262		= NewExcitation("psiRe_k262", "", "Eigenmode spatial profile")
	psiRe_k263		= NewExcitation("psiRe_k263", "", "Eigenmode spatial profile")
	psiRe_k264		= NewExcitation("psiRe_k264", "", "Eigenmode spatial profile")
	psiRe_k265		= NewExcitation("psiRe_k265", "", "Eigenmode spatial profile")
	psiRe_k266		= NewExcitation("psiRe_k266", "", "Eigenmode spatial profile")
	psiRe_k267		= NewExcitation("psiRe_k267", "", "Eigenmode spatial profile")
	psiRe_k268		= NewExcitation("psiRe_k268", "", "Eigenmode spatial profile")
	psiRe_k269		= NewExcitation("psiRe_k269", "", "Eigenmode spatial profile")
	psiRe_k270		= NewExcitation("psiRe_k270", "", "Eigenmode spatial profile")
	psiRe_k271		= NewExcitation("psiRe_k271", "", "Eigenmode spatial profile")
	psiRe_k272		= NewExcitation("psiRe_k272", "", "Eigenmode spatial profile")
	psiRe_k273		= NewExcitation("psiRe_k273", "", "Eigenmode spatial profile")
	psiRe_k274		= NewExcitation("psiRe_k274", "", "Eigenmode spatial profile")
	psiRe_k275		= NewExcitation("psiRe_k275", "", "Eigenmode spatial profile")
	psiRe_k276		= NewExcitation("psiRe_k276", "", "Eigenmode spatial profile")
	psiRe_k277		= NewExcitation("psiRe_k277", "", "Eigenmode spatial profile")
	psiRe_k278		= NewExcitation("psiRe_k278", "", "Eigenmode spatial profile")
	psiRe_k279		= NewExcitation("psiRe_k279", "", "Eigenmode spatial profile")
	psiRe_k280		= NewExcitation("psiRe_k280", "", "Eigenmode spatial profile")
	psiRe_k281		= NewExcitation("psiRe_k281", "", "Eigenmode spatial profile")
	psiRe_k282		= NewExcitation("psiRe_k282", "", "Eigenmode spatial profile")
	psiRe_k283		= NewExcitation("psiRe_k283", "", "Eigenmode spatial profile")
	psiRe_k284		= NewExcitation("psiRe_k284", "", "Eigenmode spatial profile")
	psiRe_k285		= NewExcitation("psiRe_k285", "", "Eigenmode spatial profile")
	psiRe_k286		= NewExcitation("psiRe_k286", "", "Eigenmode spatial profile")
	psiRe_k287		= NewExcitation("psiRe_k287", "", "Eigenmode spatial profile")
	psiRe_k288		= NewExcitation("psiRe_k288", "", "Eigenmode spatial profile")
	psiRe_k289		= NewExcitation("psiRe_k289", "", "Eigenmode spatial profile")
	psiRe_k290		= NewExcitation("psiRe_k290", "", "Eigenmode spatial profile")
	psiRe_k291		= NewExcitation("psiRe_k291", "", "Eigenmode spatial profile")
	psiRe_k292		= NewExcitation("psiRe_k292", "", "Eigenmode spatial profile")
	psiRe_k293		= NewExcitation("psiRe_k293", "", "Eigenmode spatial profile")
	psiRe_k294		= NewExcitation("psiRe_k294", "", "Eigenmode spatial profile")
	psiRe_k295		= NewExcitation("psiRe_k295", "", "Eigenmode spatial profile")
	psiRe_k296		= NewExcitation("psiRe_k296", "", "Eigenmode spatial profile")
	psiRe_k297		= NewExcitation("psiRe_k297", "", "Eigenmode spatial profile")
	psiRe_k298		= NewExcitation("psiRe_k298", "", "Eigenmode spatial profile")
	psiRe_k299		= NewExcitation("psiRe_k299", "", "Eigenmode spatial profile")
	psiRe_k300		= NewExcitation("psiRe_k300", "", "Eigenmode spatial profile")
	psiRe_k301		= NewExcitation("psiRe_k301", "", "Eigenmode spatial profile")
	psiRe_k302		= NewExcitation("psiRe_k302", "", "Eigenmode spatial profile")
	psiRe_k303		= NewExcitation("psiRe_k303", "", "Eigenmode spatial profile")
	psiRe_k304		= NewExcitation("psiRe_k304", "", "Eigenmode spatial profile")
	psiRe_k305		= NewExcitation("psiRe_k305", "", "Eigenmode spatial profile")
	psiRe_k306		= NewExcitation("psiRe_k306", "", "Eigenmode spatial profile")
	psiRe_k307		= NewExcitation("psiRe_k307", "", "Eigenmode spatial profile")
	psiRe_k308		= NewExcitation("psiRe_k308", "", "Eigenmode spatial profile")
	psiRe_k309		= NewExcitation("psiRe_k309", "", "Eigenmode spatial profile")
	psiRe_k310		= NewExcitation("psiRe_k310", "", "Eigenmode spatial profile")
	psiRe_k311		= NewExcitation("psiRe_k311", "", "Eigenmode spatial profile")
	psiRe_k312		= NewExcitation("psiRe_k312", "", "Eigenmode spatial profile")
	psiRe_k313		= NewExcitation("psiRe_k313", "", "Eigenmode spatial profile")
	psiRe_k314		= NewExcitation("psiRe_k314", "", "Eigenmode spatial profile")
	psiRe_k315		= NewExcitation("psiRe_k315", "", "Eigenmode spatial profile")
	psiRe_k316		= NewExcitation("psiRe_k316", "", "Eigenmode spatial profile")
	psiRe_k317		= NewExcitation("psiRe_k317", "", "Eigenmode spatial profile")
	psiRe_k318		= NewExcitation("psiRe_k318", "", "Eigenmode spatial profile")
	psiRe_k319		= NewExcitation("psiRe_k319", "", "Eigenmode spatial profile")
	psiRe_k320		= NewExcitation("psiRe_k320", "", "Eigenmode spatial profile")
	psiRe_k321		= NewExcitation("psiRe_k321", "", "Eigenmode spatial profile")
	psiRe_k322		= NewExcitation("psiRe_k322", "", "Eigenmode spatial profile")
	psiRe_k323		= NewExcitation("psiRe_k323", "", "Eigenmode spatial profile")
	psiRe_k324		= NewExcitation("psiRe_k324", "", "Eigenmode spatial profile")
	psiRe_k325		= NewExcitation("psiRe_k325", "", "Eigenmode spatial profile")
	psiRe_k326		= NewExcitation("psiRe_k326", "", "Eigenmode spatial profile")
	psiRe_k327		= NewExcitation("psiRe_k327", "", "Eigenmode spatial profile")
	psiRe_k328		= NewExcitation("psiRe_k328", "", "Eigenmode spatial profile")
	psiRe_k329		= NewExcitation("psiRe_k329", "", "Eigenmode spatial profile")
	psiRe_k330		= NewExcitation("psiRe_k330", "", "Eigenmode spatial profile")
	psiRe_k331		= NewExcitation("psiRe_k331", "", "Eigenmode spatial profile")
	psiRe_k332		= NewExcitation("psiRe_k332", "", "Eigenmode spatial profile")
	psiRe_k333		= NewExcitation("psiRe_k333", "", "Eigenmode spatial profile")
	psiRe_k334		= NewExcitation("psiRe_k334", "", "Eigenmode spatial profile")
	psiRe_k335		= NewExcitation("psiRe_k335", "", "Eigenmode spatial profile")
	psiRe_k336		= NewExcitation("psiRe_k336", "", "Eigenmode spatial profile")
	psiRe_k337		= NewExcitation("psiRe_k337", "", "Eigenmode spatial profile")
	psiRe_k338		= NewExcitation("psiRe_k338", "", "Eigenmode spatial profile")
	psiRe_k339		= NewExcitation("psiRe_k339", "", "Eigenmode spatial profile")
	psiRe_k340		= NewExcitation("psiRe_k340", "", "Eigenmode spatial profile")
	psiRe_k341		= NewExcitation("psiRe_k341", "", "Eigenmode spatial profile")
	psiRe_k342		= NewExcitation("psiRe_k342", "", "Eigenmode spatial profile")
	psiRe_k343		= NewExcitation("psiRe_k343", "", "Eigenmode spatial profile")
	psiRe_k344		= NewExcitation("psiRe_k344", "", "Eigenmode spatial profile")
	psiRe_k345		= NewExcitation("psiRe_k345", "", "Eigenmode spatial profile")
	psiRe_k346		= NewExcitation("psiRe_k346", "", "Eigenmode spatial profile")
	psiRe_k347		= NewExcitation("psiRe_k347", "", "Eigenmode spatial profile")
	psiRe_k348		= NewExcitation("psiRe_k348", "", "Eigenmode spatial profile")
	psiRe_k349		= NewExcitation("psiRe_k349", "", "Eigenmode spatial profile")
	psiRe_k350		= NewExcitation("psiRe_k350", "", "Eigenmode spatial profile")
	psiRe_k351		= NewExcitation("psiRe_k351", "", "Eigenmode spatial profile")
	psiRe_k352		= NewExcitation("psiRe_k352", "", "Eigenmode spatial profile")
	psiRe_k353		= NewExcitation("psiRe_k353", "", "Eigenmode spatial profile")
	psiRe_k354		= NewExcitation("psiRe_k354", "", "Eigenmode spatial profile")
	psiRe_k355		= NewExcitation("psiRe_k355", "", "Eigenmode spatial profile")
	psiRe_k356		= NewExcitation("psiRe_k356", "", "Eigenmode spatial profile")
	psiRe_k357		= NewExcitation("psiRe_k357", "", "Eigenmode spatial profile")
	psiRe_k358		= NewExcitation("psiRe_k358", "", "Eigenmode spatial profile")
	psiRe_k359		= NewExcitation("psiRe_k359", "", "Eigenmode spatial profile")
	psiRe_k360		= NewExcitation("psiRe_k360", "", "Eigenmode spatial profile")
	psiRe_k361		= NewExcitation("psiRe_k361", "", "Eigenmode spatial profile")
	psiRe_k362		= NewExcitation("psiRe_k362", "", "Eigenmode spatial profile")
	psiRe_k363		= NewExcitation("psiRe_k363", "", "Eigenmode spatial profile")
	psiRe_k364		= NewExcitation("psiRe_k364", "", "Eigenmode spatial profile")
	psiRe_k365		= NewExcitation("psiRe_k365", "", "Eigenmode spatial profile")
	psiRe_k366		= NewExcitation("psiRe_k366", "", "Eigenmode spatial profile")
	psiRe_k367		= NewExcitation("psiRe_k367", "", "Eigenmode spatial profile")
	psiRe_k368		= NewExcitation("psiRe_k368", "", "Eigenmode spatial profile")
	psiRe_k369		= NewExcitation("psiRe_k369", "", "Eigenmode spatial profile")
	psiRe_k370		= NewExcitation("psiRe_k370", "", "Eigenmode spatial profile")
	psiRe_k371		= NewExcitation("psiRe_k371", "", "Eigenmode spatial profile")
	psiRe_k372		= NewExcitation("psiRe_k372", "", "Eigenmode spatial profile")
	psiRe_k373		= NewExcitation("psiRe_k373", "", "Eigenmode spatial profile")
	psiRe_k374		= NewExcitation("psiRe_k374", "", "Eigenmode spatial profile")
	psiRe_k375		= NewExcitation("psiRe_k375", "", "Eigenmode spatial profile")
	psiRe_k376		= NewExcitation("psiRe_k376", "", "Eigenmode spatial profile")
	psiRe_k377		= NewExcitation("psiRe_k377", "", "Eigenmode spatial profile")
	psiRe_k378		= NewExcitation("psiRe_k378", "", "Eigenmode spatial profile")
	psiRe_k379		= NewExcitation("psiRe_k379", "", "Eigenmode spatial profile")
	psiRe_k380		= NewExcitation("psiRe_k380", "", "Eigenmode spatial profile")
	psiRe_k381		= NewExcitation("psiRe_k381", "", "Eigenmode spatial profile")
	psiRe_k382		= NewExcitation("psiRe_k382", "", "Eigenmode spatial profile")
	psiRe_k383		= NewExcitation("psiRe_k383", "", "Eigenmode spatial profile")
	psiRe_k384		= NewExcitation("psiRe_k384", "", "Eigenmode spatial profile")
	psiRe_k385		= NewExcitation("psiRe_k385", "", "Eigenmode spatial profile")
	psiRe_k386		= NewExcitation("psiRe_k386", "", "Eigenmode spatial profile")
	psiRe_k387		= NewExcitation("psiRe_k387", "", "Eigenmode spatial profile")
	psiRe_k388		= NewExcitation("psiRe_k388", "", "Eigenmode spatial profile")
	psiRe_k389		= NewExcitation("psiRe_k389", "", "Eigenmode spatial profile")
	psiRe_k390		= NewExcitation("psiRe_k390", "", "Eigenmode spatial profile")
	psiRe_k391		= NewExcitation("psiRe_k391", "", "Eigenmode spatial profile")
	psiRe_k392		= NewExcitation("psiRe_k392", "", "Eigenmode spatial profile")
	psiRe_k393		= NewExcitation("psiRe_k393", "", "Eigenmode spatial profile")
	psiRe_k394		= NewExcitation("psiRe_k394", "", "Eigenmode spatial profile")
	psiRe_k395		= NewExcitation("psiRe_k395", "", "Eigenmode spatial profile")
	psiRe_k396		= NewExcitation("psiRe_k396", "", "Eigenmode spatial profile")
	psiRe_k397		= NewExcitation("psiRe_k397", "", "Eigenmode spatial profile")
	psiRe_k398		= NewExcitation("psiRe_k398", "", "Eigenmode spatial profile")
	psiRe_k399		= NewExcitation("psiRe_k399", "", "Eigenmode spatial profile")
	psiRe_k400		= NewExcitation("psiRe_k400", "", "Eigenmode spatial profile")
	psiRe_k401		= NewExcitation("psiRe_k401", "", "Eigenmode spatial profile")
	psiRe_k402		= NewExcitation("psiRe_k402", "", "Eigenmode spatial profile")
	psiRe_k403		= NewExcitation("psiRe_k403", "", "Eigenmode spatial profile")
	psiRe_k404		= NewExcitation("psiRe_k404", "", "Eigenmode spatial profile")
	psiRe_k405		= NewExcitation("psiRe_k405", "", "Eigenmode spatial profile")
	psiRe_k406		= NewExcitation("psiRe_k406", "", "Eigenmode spatial profile")
	psiRe_k407		= NewExcitation("psiRe_k407", "", "Eigenmode spatial profile")
	psiRe_k408		= NewExcitation("psiRe_k408", "", "Eigenmode spatial profile")
	psiRe_k409		= NewExcitation("psiRe_k409", "", "Eigenmode spatial profile")
	psiRe_k410		= NewExcitation("psiRe_k410", "", "Eigenmode spatial profile")
	psiRe_k411		= NewExcitation("psiRe_k411", "", "Eigenmode spatial profile")
	psiRe_k412		= NewExcitation("psiRe_k412", "", "Eigenmode spatial profile")
	psiRe_k413		= NewExcitation("psiRe_k413", "", "Eigenmode spatial profile")
	psiRe_k414		= NewExcitation("psiRe_k414", "", "Eigenmode spatial profile")
	psiRe_k415		= NewExcitation("psiRe_k415", "", "Eigenmode spatial profile")
	psiRe_k416		= NewExcitation("psiRe_k416", "", "Eigenmode spatial profile")
	psiRe_k417		= NewExcitation("psiRe_k417", "", "Eigenmode spatial profile")
	psiRe_k418		= NewExcitation("psiRe_k418", "", "Eigenmode spatial profile")
	psiRe_k419		= NewExcitation("psiRe_k419", "", "Eigenmode spatial profile")
	psiRe_k420		= NewExcitation("psiRe_k420", "", "Eigenmode spatial profile")
	psiRe_k421		= NewExcitation("psiRe_k421", "", "Eigenmode spatial profile")
	psiRe_k422		= NewExcitation("psiRe_k422", "", "Eigenmode spatial profile")
	psiRe_k423		= NewExcitation("psiRe_k423", "", "Eigenmode spatial profile")
	psiRe_k424		= NewExcitation("psiRe_k424", "", "Eigenmode spatial profile")
	psiRe_k425		= NewExcitation("psiRe_k425", "", "Eigenmode spatial profile")
	psiRe_k426		= NewExcitation("psiRe_k426", "", "Eigenmode spatial profile")
	psiRe_k427		= NewExcitation("psiRe_k427", "", "Eigenmode spatial profile")
	psiRe_k428		= NewExcitation("psiRe_k428", "", "Eigenmode spatial profile")
	psiRe_k429		= NewExcitation("psiRe_k429", "", "Eigenmode spatial profile")
	psiRe_k430		= NewExcitation("psiRe_k430", "", "Eigenmode spatial profile")
	psiRe_k431		= NewExcitation("psiRe_k431", "", "Eigenmode spatial profile")
	psiRe_k432		= NewExcitation("psiRe_k432", "", "Eigenmode spatial profile")
	psiRe_k433		= NewExcitation("psiRe_k433", "", "Eigenmode spatial profile")
	psiRe_k434		= NewExcitation("psiRe_k434", "", "Eigenmode spatial profile")
	psiRe_k435		= NewExcitation("psiRe_k435", "", "Eigenmode spatial profile")
	psiRe_k436		= NewExcitation("psiRe_k436", "", "Eigenmode spatial profile")
	psiRe_k437		= NewExcitation("psiRe_k437", "", "Eigenmode spatial profile")
	psiRe_k438		= NewExcitation("psiRe_k438", "", "Eigenmode spatial profile")
	psiRe_k439		= NewExcitation("psiRe_k439", "", "Eigenmode spatial profile")
	psiRe_k440		= NewExcitation("psiRe_k440", "", "Eigenmode spatial profile")
	psiRe_k441		= NewExcitation("psiRe_k441", "", "Eigenmode spatial profile")
	psiRe_k442		= NewExcitation("psiRe_k442", "", "Eigenmode spatial profile")
	psiRe_k443		= NewExcitation("psiRe_k443", "", "Eigenmode spatial profile")
	psiRe_k444		= NewExcitation("psiRe_k444", "", "Eigenmode spatial profile")
	psiRe_k445		= NewExcitation("psiRe_k445", "", "Eigenmode spatial profile")
	psiRe_k446		= NewExcitation("psiRe_k446", "", "Eigenmode spatial profile")
	psiRe_k447		= NewExcitation("psiRe_k447", "", "Eigenmode spatial profile")
	psiRe_k448		= NewExcitation("psiRe_k448", "", "Eigenmode spatial profile")
	psiRe_k449		= NewExcitation("psiRe_k449", "", "Eigenmode spatial profile")
	psiRe_k450		= NewExcitation("psiRe_k450", "", "Eigenmode spatial profile")
	psiRe_k451		= NewExcitation("psiRe_k451", "", "Eigenmode spatial profile")
	psiRe_k452		= NewExcitation("psiRe_k452", "", "Eigenmode spatial profile")
	psiRe_k453		= NewExcitation("psiRe_k453", "", "Eigenmode spatial profile")
	psiRe_k454		= NewExcitation("psiRe_k454", "", "Eigenmode spatial profile")
	psiRe_k455		= NewExcitation("psiRe_k455", "", "Eigenmode spatial profile")
	psiRe_k456		= NewExcitation("psiRe_k456", "", "Eigenmode spatial profile")
	psiRe_k457		= NewExcitation("psiRe_k457", "", "Eigenmode spatial profile")
	psiRe_k458		= NewExcitation("psiRe_k458", "", "Eigenmode spatial profile")
	psiRe_k459		= NewExcitation("psiRe_k459", "", "Eigenmode spatial profile")
	psiRe_k460		= NewExcitation("psiRe_k460", "", "Eigenmode spatial profile")
	psiRe_k461		= NewExcitation("psiRe_k461", "", "Eigenmode spatial profile")
	psiRe_k462		= NewExcitation("psiRe_k462", "", "Eigenmode spatial profile")
	psiRe_k463		= NewExcitation("psiRe_k463", "", "Eigenmode spatial profile")
	psiRe_k464		= NewExcitation("psiRe_k464", "", "Eigenmode spatial profile")
	psiRe_k465		= NewExcitation("psiRe_k465", "", "Eigenmode spatial profile")
	psiRe_k466		= NewExcitation("psiRe_k466", "", "Eigenmode spatial profile")
	psiRe_k467		= NewExcitation("psiRe_k467", "", "Eigenmode spatial profile")
	psiRe_k468		= NewExcitation("psiRe_k468", "", "Eigenmode spatial profile")
	psiRe_k469		= NewExcitation("psiRe_k469", "", "Eigenmode spatial profile")
	psiRe_k470		= NewExcitation("psiRe_k470", "", "Eigenmode spatial profile")
	psiRe_k471		= NewExcitation("psiRe_k471", "", "Eigenmode spatial profile")
	psiRe_k472		= NewExcitation("psiRe_k472", "", "Eigenmode spatial profile")
	psiRe_k473		= NewExcitation("psiRe_k473", "", "Eigenmode spatial profile")
	psiRe_k474		= NewExcitation("psiRe_k474", "", "Eigenmode spatial profile")
	psiRe_k475		= NewExcitation("psiRe_k475", "", "Eigenmode spatial profile")
	psiRe_k476		= NewExcitation("psiRe_k476", "", "Eigenmode spatial profile")
	psiRe_k477		= NewExcitation("psiRe_k477", "", "Eigenmode spatial profile")
	psiRe_k478		= NewExcitation("psiRe_k478", "", "Eigenmode spatial profile")
	psiRe_k479		= NewExcitation("psiRe_k479", "", "Eigenmode spatial profile")
	psiRe_k480		= NewExcitation("psiRe_k480", "", "Eigenmode spatial profile")
	psiRe_k481		= NewExcitation("psiRe_k481", "", "Eigenmode spatial profile")
	psiRe_k482		= NewExcitation("psiRe_k482", "", "Eigenmode spatial profile")
	psiRe_k483		= NewExcitation("psiRe_k483", "", "Eigenmode spatial profile")
	psiRe_k484		= NewExcitation("psiRe_k484", "", "Eigenmode spatial profile")
	psiRe_k485		= NewExcitation("psiRe_k485", "", "Eigenmode spatial profile")
	psiRe_k486		= NewExcitation("psiRe_k486", "", "Eigenmode spatial profile")
	psiRe_k487		= NewExcitation("psiRe_k487", "", "Eigenmode spatial profile")
	psiRe_k488		= NewExcitation("psiRe_k488", "", "Eigenmode spatial profile")
	psiRe_k489		= NewExcitation("psiRe_k489", "", "Eigenmode spatial profile")
	psiRe_k490		= NewExcitation("psiRe_k490", "", "Eigenmode spatial profile")
	psiRe_k491		= NewExcitation("psiRe_k491", "", "Eigenmode spatial profile")
	psiRe_k492		= NewExcitation("psiRe_k492", "", "Eigenmode spatial profile")
	psiRe_k493		= NewExcitation("psiRe_k493", "", "Eigenmode spatial profile")
	psiRe_k494		= NewExcitation("psiRe_k494", "", "Eigenmode spatial profile")
	psiRe_k495		= NewExcitation("psiRe_k495", "", "Eigenmode spatial profile")
	psiRe_k496		= NewExcitation("psiRe_k496", "", "Eigenmode spatial profile")
	psiRe_k497		= NewExcitation("psiRe_k497", "", "Eigenmode spatial profile")
	psiRe_k498		= NewExcitation("psiRe_k498", "", "Eigenmode spatial profile")
	psiRe_k499		= NewExcitation("psiRe_k499", "", "Eigenmode spatial profile")
	
	psiIm_k000		= NewExcitation("psiIm_k000", "", "Eigenmode spatial profile")
	psiIm_k001		= NewExcitation("psiIm_k001", "", "Eigenmode spatial profile")
	psiIm_k002		= NewExcitation("psiIm_k002", "", "Eigenmode spatial profile")
	psiIm_k003		= NewExcitation("psiIm_k003", "", "Eigenmode spatial profile")
	psiIm_k004		= NewExcitation("psiIm_k004", "", "Eigenmode spatial profile")
	psiIm_k005		= NewExcitation("psiIm_k005", "", "Eigenmode spatial profile")
	psiIm_k006		= NewExcitation("psiIm_k006", "", "Eigenmode spatial profile")
	psiIm_k007		= NewExcitation("psiIm_k007", "", "Eigenmode spatial profile")
	psiIm_k008		= NewExcitation("psiIm_k008", "", "Eigenmode spatial profile")
	psiIm_k009		= NewExcitation("psiIm_k009", "", "Eigenmode spatial profile")
	psiIm_k010		= NewExcitation("psiIm_k010", "", "Eigenmode spatial profile")
	psiIm_k011		= NewExcitation("psiIm_k011", "", "Eigenmode spatial profile")
	psiIm_k012		= NewExcitation("psiIm_k012", "", "Eigenmode spatial profile")
	psiIm_k013		= NewExcitation("psiIm_k013", "", "Eigenmode spatial profile")
	psiIm_k014		= NewExcitation("psiIm_k014", "", "Eigenmode spatial profile")
	psiIm_k015		= NewExcitation("psiIm_k015", "", "Eigenmode spatial profile")
	psiIm_k016		= NewExcitation("psiIm_k016", "", "Eigenmode spatial profile")
	psiIm_k017		= NewExcitation("psiIm_k017", "", "Eigenmode spatial profile")
	psiIm_k018		= NewExcitation("psiIm_k018", "", "Eigenmode spatial profile")
	psiIm_k019		= NewExcitation("psiIm_k019", "", "Eigenmode spatial profile")
	psiIm_k020		= NewExcitation("psiIm_k020", "", "Eigenmode spatial profile")
	psiIm_k021		= NewExcitation("psiIm_k021", "", "Eigenmode spatial profile")
	psiIm_k022		= NewExcitation("psiIm_k022", "", "Eigenmode spatial profile")
	psiIm_k023		= NewExcitation("psiIm_k023", "", "Eigenmode spatial profile")
	psiIm_k024		= NewExcitation("psiIm_k024", "", "Eigenmode spatial profile")
	psiIm_k025		= NewExcitation("psiIm_k025", "", "Eigenmode spatial profile")
	psiIm_k026		= NewExcitation("psiIm_k026", "", "Eigenmode spatial profile")
	psiIm_k027		= NewExcitation("psiIm_k027", "", "Eigenmode spatial profile")
	psiIm_k028		= NewExcitation("psiIm_k028", "", "Eigenmode spatial profile")
	psiIm_k029		= NewExcitation("psiIm_k029", "", "Eigenmode spatial profile")
	psiIm_k030		= NewExcitation("psiIm_k030", "", "Eigenmode spatial profile")
	psiIm_k031		= NewExcitation("psiIm_k031", "", "Eigenmode spatial profile")
	psiIm_k032		= NewExcitation("psiIm_k032", "", "Eigenmode spatial profile")
	psiIm_k033		= NewExcitation("psiIm_k033", "", "Eigenmode spatial profile")
	psiIm_k034		= NewExcitation("psiIm_k034", "", "Eigenmode spatial profile")
	psiIm_k035		= NewExcitation("psiIm_k035", "", "Eigenmode spatial profile")
	psiIm_k036		= NewExcitation("psiIm_k036", "", "Eigenmode spatial profile")
	psiIm_k037		= NewExcitation("psiIm_k037", "", "Eigenmode spatial profile")
	psiIm_k038		= NewExcitation("psiIm_k038", "", "Eigenmode spatial profile")
	psiIm_k039		= NewExcitation("psiIm_k039", "", "Eigenmode spatial profile")
	psiIm_k040		= NewExcitation("psiIm_k040", "", "Eigenmode spatial profile")
	psiIm_k041		= NewExcitation("psiIm_k041", "", "Eigenmode spatial profile")
	psiIm_k042		= NewExcitation("psiIm_k042", "", "Eigenmode spatial profile")
	psiIm_k043		= NewExcitation("psiIm_k043", "", "Eigenmode spatial profile")
	psiIm_k044		= NewExcitation("psiIm_k044", "", "Eigenmode spatial profile")
	psiIm_k045		= NewExcitation("psiIm_k045", "", "Eigenmode spatial profile")
	psiIm_k046		= NewExcitation("psiIm_k046", "", "Eigenmode spatial profile")
	psiIm_k047		= NewExcitation("psiIm_k047", "", "Eigenmode spatial profile")
	psiIm_k048		= NewExcitation("psiIm_k048", "", "Eigenmode spatial profile")
	psiIm_k049		= NewExcitation("psiIm_k049", "", "Eigenmode spatial profile")
	psiIm_k050		= NewExcitation("psiIm_k050", "", "Eigenmode spatial profile")
	psiIm_k051		= NewExcitation("psiIm_k051", "", "Eigenmode spatial profile")
	psiIm_k052		= NewExcitation("psiIm_k052", "", "Eigenmode spatial profile")
	psiIm_k053		= NewExcitation("psiIm_k053", "", "Eigenmode spatial profile")
	psiIm_k054		= NewExcitation("psiIm_k054", "", "Eigenmode spatial profile")
	psiIm_k055		= NewExcitation("psiIm_k055", "", "Eigenmode spatial profile")
	psiIm_k056		= NewExcitation("psiIm_k056", "", "Eigenmode spatial profile")
	psiIm_k057		= NewExcitation("psiIm_k057", "", "Eigenmode spatial profile")
	psiIm_k058		= NewExcitation("psiIm_k058", "", "Eigenmode spatial profile")
	psiIm_k059		= NewExcitation("psiIm_k059", "", "Eigenmode spatial profile")
	psiIm_k060		= NewExcitation("psiIm_k060", "", "Eigenmode spatial profile")
	psiIm_k061		= NewExcitation("psiIm_k061", "", "Eigenmode spatial profile")
	psiIm_k062		= NewExcitation("psiIm_k062", "", "Eigenmode spatial profile")
	psiIm_k063		= NewExcitation("psiIm_k063", "", "Eigenmode spatial profile")
	psiIm_k064		= NewExcitation("psiIm_k064", "", "Eigenmode spatial profile")
	psiIm_k065		= NewExcitation("psiIm_k065", "", "Eigenmode spatial profile")
	psiIm_k066		= NewExcitation("psiIm_k066", "", "Eigenmode spatial profile")
	psiIm_k067		= NewExcitation("psiIm_k067", "", "Eigenmode spatial profile")
	psiIm_k068		= NewExcitation("psiIm_k068", "", "Eigenmode spatial profile")
	psiIm_k069		= NewExcitation("psiIm_k069", "", "Eigenmode spatial profile")
	psiIm_k070		= NewExcitation("psiIm_k070", "", "Eigenmode spatial profile")
	psiIm_k071		= NewExcitation("psiIm_k071", "", "Eigenmode spatial profile")
	psiIm_k072		= NewExcitation("psiIm_k072", "", "Eigenmode spatial profile")
	psiIm_k073		= NewExcitation("psiIm_k073", "", "Eigenmode spatial profile")
	psiIm_k074		= NewExcitation("psiIm_k074", "", "Eigenmode spatial profile")
	psiIm_k075		= NewExcitation("psiIm_k075", "", "Eigenmode spatial profile")
	psiIm_k076		= NewExcitation("psiIm_k076", "", "Eigenmode spatial profile")
	psiIm_k077		= NewExcitation("psiIm_k077", "", "Eigenmode spatial profile")
	psiIm_k078		= NewExcitation("psiIm_k078", "", "Eigenmode spatial profile")
	psiIm_k079		= NewExcitation("psiIm_k079", "", "Eigenmode spatial profile")
	psiIm_k080		= NewExcitation("psiIm_k080", "", "Eigenmode spatial profile")
	psiIm_k081		= NewExcitation("psiIm_k081", "", "Eigenmode spatial profile")
	psiIm_k082		= NewExcitation("psiIm_k082", "", "Eigenmode spatial profile")
	psiIm_k083		= NewExcitation("psiIm_k083", "", "Eigenmode spatial profile")
	psiIm_k084		= NewExcitation("psiIm_k084", "", "Eigenmode spatial profile")
	psiIm_k085		= NewExcitation("psiIm_k085", "", "Eigenmode spatial profile")
	psiIm_k086		= NewExcitation("psiIm_k086", "", "Eigenmode spatial profile")
	psiIm_k087		= NewExcitation("psiIm_k087", "", "Eigenmode spatial profile")
	psiIm_k088		= NewExcitation("psiIm_k088", "", "Eigenmode spatial profile")
	psiIm_k089		= NewExcitation("psiIm_k089", "", "Eigenmode spatial profile")
	psiIm_k090		= NewExcitation("psiIm_k090", "", "Eigenmode spatial profile")
	psiIm_k091		= NewExcitation("psiIm_k091", "", "Eigenmode spatial profile")
	psiIm_k092		= NewExcitation("psiIm_k092", "", "Eigenmode spatial profile")
	psiIm_k093		= NewExcitation("psiIm_k093", "", "Eigenmode spatial profile")
	psiIm_k094		= NewExcitation("psiIm_k094", "", "Eigenmode spatial profile")
	psiIm_k095		= NewExcitation("psiIm_k095", "", "Eigenmode spatial profile")
	psiIm_k096		= NewExcitation("psiIm_k096", "", "Eigenmode spatial profile")
	psiIm_k097		= NewExcitation("psiIm_k097", "", "Eigenmode spatial profile")
	psiIm_k098		= NewExcitation("psiIm_k098", "", "Eigenmode spatial profile")
	psiIm_k099		= NewExcitation("psiIm_k099", "", "Eigenmode spatial profile")
	psiIm_k100		= NewExcitation("psiIm_k100", "", "Eigenmode spatial profile")
	psiIm_k101		= NewExcitation("psiIm_k101", "", "Eigenmode spatial profile")
	psiIm_k102		= NewExcitation("psiIm_k102", "", "Eigenmode spatial profile")
	psiIm_k103		= NewExcitation("psiIm_k103", "", "Eigenmode spatial profile")
	psiIm_k104		= NewExcitation("psiIm_k104", "", "Eigenmode spatial profile")
	psiIm_k105		= NewExcitation("psiIm_k105", "", "Eigenmode spatial profile")
	psiIm_k106		= NewExcitation("psiIm_k106", "", "Eigenmode spatial profile")
	psiIm_k107		= NewExcitation("psiIm_k107", "", "Eigenmode spatial profile")
	psiIm_k108		= NewExcitation("psiIm_k108", "", "Eigenmode spatial profile")
	psiIm_k109		= NewExcitation("psiIm_k109", "", "Eigenmode spatial profile")
	psiIm_k110		= NewExcitation("psiIm_k110", "", "Eigenmode spatial profile")
	psiIm_k111		= NewExcitation("psiIm_k111", "", "Eigenmode spatial profile")
	psiIm_k112		= NewExcitation("psiIm_k112", "", "Eigenmode spatial profile")
	psiIm_k113		= NewExcitation("psiIm_k113", "", "Eigenmode spatial profile")
	psiIm_k114		= NewExcitation("psiIm_k114", "", "Eigenmode spatial profile")
	psiIm_k115		= NewExcitation("psiIm_k115", "", "Eigenmode spatial profile")
	psiIm_k116		= NewExcitation("psiIm_k116", "", "Eigenmode spatial profile")
	psiIm_k117		= NewExcitation("psiIm_k117", "", "Eigenmode spatial profile")
	psiIm_k118		= NewExcitation("psiIm_k118", "", "Eigenmode spatial profile")
	psiIm_k119		= NewExcitation("psiIm_k119", "", "Eigenmode spatial profile")
	psiIm_k120		= NewExcitation("psiIm_k120", "", "Eigenmode spatial profile")
	psiIm_k121		= NewExcitation("psiIm_k121", "", "Eigenmode spatial profile")
	psiIm_k122		= NewExcitation("psiIm_k122", "", "Eigenmode spatial profile")
	psiIm_k123		= NewExcitation("psiIm_k123", "", "Eigenmode spatial profile")
	psiIm_k124		= NewExcitation("psiIm_k124", "", "Eigenmode spatial profile")
	psiIm_k125		= NewExcitation("psiIm_k125", "", "Eigenmode spatial profile")
	psiIm_k126		= NewExcitation("psiIm_k126", "", "Eigenmode spatial profile")
	psiIm_k127		= NewExcitation("psiIm_k127", "", "Eigenmode spatial profile")
	psiIm_k128		= NewExcitation("psiIm_k128", "", "Eigenmode spatial profile")
	psiIm_k129		= NewExcitation("psiIm_k129", "", "Eigenmode spatial profile")
	psiIm_k130		= NewExcitation("psiIm_k130", "", "Eigenmode spatial profile")
	psiIm_k131		= NewExcitation("psiIm_k131", "", "Eigenmode spatial profile")
	psiIm_k132		= NewExcitation("psiIm_k132", "", "Eigenmode spatial profile")
	psiIm_k133		= NewExcitation("psiIm_k133", "", "Eigenmode spatial profile")
	psiIm_k134		= NewExcitation("psiIm_k134", "", "Eigenmode spatial profile")
	psiIm_k135		= NewExcitation("psiIm_k135", "", "Eigenmode spatial profile")
	psiIm_k136		= NewExcitation("psiIm_k136", "", "Eigenmode spatial profile")
	psiIm_k137		= NewExcitation("psiIm_k137", "", "Eigenmode spatial profile")
	psiIm_k138		= NewExcitation("psiIm_k138", "", "Eigenmode spatial profile")
	psiIm_k139		= NewExcitation("psiIm_k139", "", "Eigenmode spatial profile")
	psiIm_k140		= NewExcitation("psiIm_k140", "", "Eigenmode spatial profile")
	psiIm_k141		= NewExcitation("psiIm_k141", "", "Eigenmode spatial profile")
	psiIm_k142		= NewExcitation("psiIm_k142", "", "Eigenmode spatial profile")
	psiIm_k143		= NewExcitation("psiIm_k143", "", "Eigenmode spatial profile")
	psiIm_k144		= NewExcitation("psiIm_k144", "", "Eigenmode spatial profile")
	psiIm_k145		= NewExcitation("psiIm_k145", "", "Eigenmode spatial profile")
	psiIm_k146		= NewExcitation("psiIm_k146", "", "Eigenmode spatial profile")
	psiIm_k147		= NewExcitation("psiIm_k147", "", "Eigenmode spatial profile")
	psiIm_k148		= NewExcitation("psiIm_k148", "", "Eigenmode spatial profile")
	psiIm_k149		= NewExcitation("psiIm_k149", "", "Eigenmode spatial profile")
	psiIm_k150		= NewExcitation("psiIm_k150", "", "Eigenmode spatial profile")
	psiIm_k151		= NewExcitation("psiIm_k151", "", "Eigenmode spatial profile")
	psiIm_k152		= NewExcitation("psiIm_k152", "", "Eigenmode spatial profile")
	psiIm_k153		= NewExcitation("psiIm_k153", "", "Eigenmode spatial profile")
	psiIm_k154		= NewExcitation("psiIm_k154", "", "Eigenmode spatial profile")
	psiIm_k155		= NewExcitation("psiIm_k155", "", "Eigenmode spatial profile")
	psiIm_k156		= NewExcitation("psiIm_k156", "", "Eigenmode spatial profile")
	psiIm_k157		= NewExcitation("psiIm_k157", "", "Eigenmode spatial profile")
	psiIm_k158		= NewExcitation("psiIm_k158", "", "Eigenmode spatial profile")
	psiIm_k159		= NewExcitation("psiIm_k159", "", "Eigenmode spatial profile")
	psiIm_k160		= NewExcitation("psiIm_k160", "", "Eigenmode spatial profile")
	psiIm_k161		= NewExcitation("psiIm_k161", "", "Eigenmode spatial profile")
	psiIm_k162		= NewExcitation("psiIm_k162", "", "Eigenmode spatial profile")
	psiIm_k163		= NewExcitation("psiIm_k163", "", "Eigenmode spatial profile")
	psiIm_k164		= NewExcitation("psiIm_k164", "", "Eigenmode spatial profile")
	psiIm_k165		= NewExcitation("psiIm_k165", "", "Eigenmode spatial profile")
	psiIm_k166		= NewExcitation("psiIm_k166", "", "Eigenmode spatial profile")
	psiIm_k167		= NewExcitation("psiIm_k167", "", "Eigenmode spatial profile")
	psiIm_k168		= NewExcitation("psiIm_k168", "", "Eigenmode spatial profile")
	psiIm_k169		= NewExcitation("psiIm_k169", "", "Eigenmode spatial profile")
	psiIm_k170		= NewExcitation("psiIm_k170", "", "Eigenmode spatial profile")
	psiIm_k171		= NewExcitation("psiIm_k171", "", "Eigenmode spatial profile")
	psiIm_k172		= NewExcitation("psiIm_k172", "", "Eigenmode spatial profile")
	psiIm_k173		= NewExcitation("psiIm_k173", "", "Eigenmode spatial profile")
	psiIm_k174		= NewExcitation("psiIm_k174", "", "Eigenmode spatial profile")
	psiIm_k175		= NewExcitation("psiIm_k175", "", "Eigenmode spatial profile")
	psiIm_k176		= NewExcitation("psiIm_k176", "", "Eigenmode spatial profile")
	psiIm_k177		= NewExcitation("psiIm_k177", "", "Eigenmode spatial profile")
	psiIm_k178		= NewExcitation("psiIm_k178", "", "Eigenmode spatial profile")
	psiIm_k179		= NewExcitation("psiIm_k179", "", "Eigenmode spatial profile")
	psiIm_k180		= NewExcitation("psiIm_k180", "", "Eigenmode spatial profile")
	psiIm_k181		= NewExcitation("psiIm_k181", "", "Eigenmode spatial profile")
	psiIm_k182		= NewExcitation("psiIm_k182", "", "Eigenmode spatial profile")
	psiIm_k183		= NewExcitation("psiIm_k183", "", "Eigenmode spatial profile")
	psiIm_k184		= NewExcitation("psiIm_k184", "", "Eigenmode spatial profile")
	psiIm_k185		= NewExcitation("psiIm_k185", "", "Eigenmode spatial profile")
	psiIm_k186		= NewExcitation("psiIm_k186", "", "Eigenmode spatial profile")
	psiIm_k187		= NewExcitation("psiIm_k187", "", "Eigenmode spatial profile")
	psiIm_k188		= NewExcitation("psiIm_k188", "", "Eigenmode spatial profile")
	psiIm_k189		= NewExcitation("psiIm_k189", "", "Eigenmode spatial profile")
	psiIm_k190		= NewExcitation("psiIm_k190", "", "Eigenmode spatial profile")
	psiIm_k191		= NewExcitation("psiIm_k191", "", "Eigenmode spatial profile")
	psiIm_k192		= NewExcitation("psiIm_k192", "", "Eigenmode spatial profile")
	psiIm_k193		= NewExcitation("psiIm_k193", "", "Eigenmode spatial profile")
	psiIm_k194		= NewExcitation("psiIm_k194", "", "Eigenmode spatial profile")
	psiIm_k195		= NewExcitation("psiIm_k195", "", "Eigenmode spatial profile")
	psiIm_k196		= NewExcitation("psiIm_k196", "", "Eigenmode spatial profile")
	psiIm_k197		= NewExcitation("psiIm_k197", "", "Eigenmode spatial profile")
	psiIm_k198		= NewExcitation("psiIm_k198", "", "Eigenmode spatial profile")
	psiIm_k199		= NewExcitation("psiIm_k199", "", "Eigenmode spatial profile")
	psiIm_k200		= NewExcitation("psiIm_k200", "", "Eigenmode spatial profile")
	psiIm_k201		= NewExcitation("psiIm_k201", "", "Eigenmode spatial profile")
	psiIm_k202		= NewExcitation("psiIm_k202", "", "Eigenmode spatial profile")
	psiIm_k203		= NewExcitation("psiIm_k203", "", "Eigenmode spatial profile")
	psiIm_k204		= NewExcitation("psiIm_k204", "", "Eigenmode spatial profile")
	psiIm_k205		= NewExcitation("psiIm_k205", "", "Eigenmode spatial profile")
	psiIm_k206		= NewExcitation("psiIm_k206", "", "Eigenmode spatial profile")
	psiIm_k207		= NewExcitation("psiIm_k207", "", "Eigenmode spatial profile")
	psiIm_k208		= NewExcitation("psiIm_k208", "", "Eigenmode spatial profile")
	psiIm_k209		= NewExcitation("psiIm_k209", "", "Eigenmode spatial profile")
	psiIm_k210		= NewExcitation("psiIm_k210", "", "Eigenmode spatial profile")
	psiIm_k211		= NewExcitation("psiIm_k211", "", "Eigenmode spatial profile")
	psiIm_k212		= NewExcitation("psiIm_k212", "", "Eigenmode spatial profile")
	psiIm_k213		= NewExcitation("psiIm_k213", "", "Eigenmode spatial profile")
	psiIm_k214		= NewExcitation("psiIm_k214", "", "Eigenmode spatial profile")
	psiIm_k215		= NewExcitation("psiIm_k215", "", "Eigenmode spatial profile")
	psiIm_k216		= NewExcitation("psiIm_k216", "", "Eigenmode spatial profile")
	psiIm_k217		= NewExcitation("psiIm_k217", "", "Eigenmode spatial profile")
	psiIm_k218		= NewExcitation("psiIm_k218", "", "Eigenmode spatial profile")
	psiIm_k219		= NewExcitation("psiIm_k219", "", "Eigenmode spatial profile")
	psiIm_k220		= NewExcitation("psiIm_k220", "", "Eigenmode spatial profile")
	psiIm_k221		= NewExcitation("psiIm_k221", "", "Eigenmode spatial profile")
	psiIm_k222		= NewExcitation("psiIm_k222", "", "Eigenmode spatial profile")
	psiIm_k223		= NewExcitation("psiIm_k223", "", "Eigenmode spatial profile")
	psiIm_k224		= NewExcitation("psiIm_k224", "", "Eigenmode spatial profile")
	psiIm_k225		= NewExcitation("psiIm_k225", "", "Eigenmode spatial profile")
	psiIm_k226		= NewExcitation("psiIm_k226", "", "Eigenmode spatial profile")
	psiIm_k227		= NewExcitation("psiIm_k227", "", "Eigenmode spatial profile")
	psiIm_k228		= NewExcitation("psiIm_k228", "", "Eigenmode spatial profile")
	psiIm_k229		= NewExcitation("psiIm_k229", "", "Eigenmode spatial profile")
	psiIm_k230		= NewExcitation("psiIm_k230", "", "Eigenmode spatial profile")
	psiIm_k231		= NewExcitation("psiIm_k231", "", "Eigenmode spatial profile")
	psiIm_k232		= NewExcitation("psiIm_k232", "", "Eigenmode spatial profile")
	psiIm_k233		= NewExcitation("psiIm_k233", "", "Eigenmode spatial profile")
	psiIm_k234		= NewExcitation("psiIm_k234", "", "Eigenmode spatial profile")
	psiIm_k235		= NewExcitation("psiIm_k235", "", "Eigenmode spatial profile")
	psiIm_k236		= NewExcitation("psiIm_k236", "", "Eigenmode spatial profile")
	psiIm_k237		= NewExcitation("psiIm_k237", "", "Eigenmode spatial profile")
	psiIm_k238		= NewExcitation("psiIm_k238", "", "Eigenmode spatial profile")
	psiIm_k239		= NewExcitation("psiIm_k239", "", "Eigenmode spatial profile")
	psiIm_k240		= NewExcitation("psiIm_k240", "", "Eigenmode spatial profile")
	psiIm_k241		= NewExcitation("psiIm_k241", "", "Eigenmode spatial profile")
	psiIm_k242		= NewExcitation("psiIm_k242", "", "Eigenmode spatial profile")
	psiIm_k243		= NewExcitation("psiIm_k243", "", "Eigenmode spatial profile")
	psiIm_k244		= NewExcitation("psiIm_k244", "", "Eigenmode spatial profile")
	psiIm_k245		= NewExcitation("psiIm_k245", "", "Eigenmode spatial profile")
	psiIm_k246		= NewExcitation("psiIm_k246", "", "Eigenmode spatial profile")
	psiIm_k247		= NewExcitation("psiIm_k247", "", "Eigenmode spatial profile")
	psiIm_k248		= NewExcitation("psiIm_k248", "", "Eigenmode spatial profile")
	psiIm_k249		= NewExcitation("psiIm_k249", "", "Eigenmode spatial profile")
	psiIm_k250		= NewExcitation("psiIm_k250", "", "Eigenmode spatial profile")
	psiIm_k251		= NewExcitation("psiIm_k251", "", "Eigenmode spatial profile")
	psiIm_k252		= NewExcitation("psiIm_k252", "", "Eigenmode spatial profile")
	psiIm_k253		= NewExcitation("psiIm_k253", "", "Eigenmode spatial profile")
	psiIm_k254		= NewExcitation("psiIm_k254", "", "Eigenmode spatial profile")
	psiIm_k255		= NewExcitation("psiIm_k255", "", "Eigenmode spatial profile")
	psiIm_k256		= NewExcitation("psiIm_k256", "", "Eigenmode spatial profile")
	psiIm_k257		= NewExcitation("psiIm_k257", "", "Eigenmode spatial profile")
	psiIm_k258		= NewExcitation("psiIm_k258", "", "Eigenmode spatial profile")
	psiIm_k259		= NewExcitation("psiIm_k259", "", "Eigenmode spatial profile")
	psiIm_k260		= NewExcitation("psiIm_k260", "", "Eigenmode spatial profile")
	psiIm_k261		= NewExcitation("psiIm_k261", "", "Eigenmode spatial profile")
	psiIm_k262		= NewExcitation("psiIm_k262", "", "Eigenmode spatial profile")
	psiIm_k263		= NewExcitation("psiIm_k263", "", "Eigenmode spatial profile")
	psiIm_k264		= NewExcitation("psiIm_k264", "", "Eigenmode spatial profile")
	psiIm_k265		= NewExcitation("psiIm_k265", "", "Eigenmode spatial profile")
	psiIm_k266		= NewExcitation("psiIm_k266", "", "Eigenmode spatial profile")
	psiIm_k267		= NewExcitation("psiIm_k267", "", "Eigenmode spatial profile")
	psiIm_k268		= NewExcitation("psiIm_k268", "", "Eigenmode spatial profile")
	psiIm_k269		= NewExcitation("psiIm_k269", "", "Eigenmode spatial profile")
	psiIm_k270		= NewExcitation("psiIm_k270", "", "Eigenmode spatial profile")
	psiIm_k271		= NewExcitation("psiIm_k271", "", "Eigenmode spatial profile")
	psiIm_k272		= NewExcitation("psiIm_k272", "", "Eigenmode spatial profile")
	psiIm_k273		= NewExcitation("psiIm_k273", "", "Eigenmode spatial profile")
	psiIm_k274		= NewExcitation("psiIm_k274", "", "Eigenmode spatial profile")
	psiIm_k275		= NewExcitation("psiIm_k275", "", "Eigenmode spatial profile")
	psiIm_k276		= NewExcitation("psiIm_k276", "", "Eigenmode spatial profile")
	psiIm_k277		= NewExcitation("psiIm_k277", "", "Eigenmode spatial profile")
	psiIm_k278		= NewExcitation("psiIm_k278", "", "Eigenmode spatial profile")
	psiIm_k279		= NewExcitation("psiIm_k279", "", "Eigenmode spatial profile")
	psiIm_k280		= NewExcitation("psiIm_k280", "", "Eigenmode spatial profile")
	psiIm_k281		= NewExcitation("psiIm_k281", "", "Eigenmode spatial profile")
	psiIm_k282		= NewExcitation("psiIm_k282", "", "Eigenmode spatial profile")
	psiIm_k283		= NewExcitation("psiIm_k283", "", "Eigenmode spatial profile")
	psiIm_k284		= NewExcitation("psiIm_k284", "", "Eigenmode spatial profile")
	psiIm_k285		= NewExcitation("psiIm_k285", "", "Eigenmode spatial profile")
	psiIm_k286		= NewExcitation("psiIm_k286", "", "Eigenmode spatial profile")
	psiIm_k287		= NewExcitation("psiIm_k287", "", "Eigenmode spatial profile")
	psiIm_k288		= NewExcitation("psiIm_k288", "", "Eigenmode spatial profile")
	psiIm_k289		= NewExcitation("psiIm_k289", "", "Eigenmode spatial profile")
	psiIm_k290		= NewExcitation("psiIm_k290", "", "Eigenmode spatial profile")
	psiIm_k291		= NewExcitation("psiIm_k291", "", "Eigenmode spatial profile")
	psiIm_k292		= NewExcitation("psiIm_k292", "", "Eigenmode spatial profile")
	psiIm_k293		= NewExcitation("psiIm_k293", "", "Eigenmode spatial profile")
	psiIm_k294		= NewExcitation("psiIm_k294", "", "Eigenmode spatial profile")
	psiIm_k295		= NewExcitation("psiIm_k295", "", "Eigenmode spatial profile")
	psiIm_k296		= NewExcitation("psiIm_k296", "", "Eigenmode spatial profile")
	psiIm_k297		= NewExcitation("psiIm_k297", "", "Eigenmode spatial profile")
	psiIm_k298		= NewExcitation("psiIm_k298", "", "Eigenmode spatial profile")
	psiIm_k299		= NewExcitation("psiIm_k299", "", "Eigenmode spatial profile")
	psiIm_k300		= NewExcitation("psiIm_k300", "", "Eigenmode spatial profile")
	psiIm_k301		= NewExcitation("psiIm_k301", "", "Eigenmode spatial profile")
	psiIm_k302		= NewExcitation("psiIm_k302", "", "Eigenmode spatial profile")
	psiIm_k303		= NewExcitation("psiIm_k303", "", "Eigenmode spatial profile")
	psiIm_k304		= NewExcitation("psiIm_k304", "", "Eigenmode spatial profile")
	psiIm_k305		= NewExcitation("psiIm_k305", "", "Eigenmode spatial profile")
	psiIm_k306		= NewExcitation("psiIm_k306", "", "Eigenmode spatial profile")
	psiIm_k307		= NewExcitation("psiIm_k307", "", "Eigenmode spatial profile")
	psiIm_k308		= NewExcitation("psiIm_k308", "", "Eigenmode spatial profile")
	psiIm_k309		= NewExcitation("psiIm_k309", "", "Eigenmode spatial profile")
	psiIm_k310		= NewExcitation("psiIm_k310", "", "Eigenmode spatial profile")
	psiIm_k311		= NewExcitation("psiIm_k311", "", "Eigenmode spatial profile")
	psiIm_k312		= NewExcitation("psiIm_k312", "", "Eigenmode spatial profile")
	psiIm_k313		= NewExcitation("psiIm_k313", "", "Eigenmode spatial profile")
	psiIm_k314		= NewExcitation("psiIm_k314", "", "Eigenmode spatial profile")
	psiIm_k315		= NewExcitation("psiIm_k315", "", "Eigenmode spatial profile")
	psiIm_k316		= NewExcitation("psiIm_k316", "", "Eigenmode spatial profile")
	psiIm_k317		= NewExcitation("psiIm_k317", "", "Eigenmode spatial profile")
	psiIm_k318		= NewExcitation("psiIm_k318", "", "Eigenmode spatial profile")
	psiIm_k319		= NewExcitation("psiIm_k319", "", "Eigenmode spatial profile")
	psiIm_k320		= NewExcitation("psiIm_k320", "", "Eigenmode spatial profile")
	psiIm_k321		= NewExcitation("psiIm_k321", "", "Eigenmode spatial profile")
	psiIm_k322		= NewExcitation("psiIm_k322", "", "Eigenmode spatial profile")
	psiIm_k323		= NewExcitation("psiIm_k323", "", "Eigenmode spatial profile")
	psiIm_k324		= NewExcitation("psiIm_k324", "", "Eigenmode spatial profile")
	psiIm_k325		= NewExcitation("psiIm_k325", "", "Eigenmode spatial profile")
	psiIm_k326		= NewExcitation("psiIm_k326", "", "Eigenmode spatial profile")
	psiIm_k327		= NewExcitation("psiIm_k327", "", "Eigenmode spatial profile")
	psiIm_k328		= NewExcitation("psiIm_k328", "", "Eigenmode spatial profile")
	psiIm_k329		= NewExcitation("psiIm_k329", "", "Eigenmode spatial profile")
	psiIm_k330		= NewExcitation("psiIm_k330", "", "Eigenmode spatial profile")
	psiIm_k331		= NewExcitation("psiIm_k331", "", "Eigenmode spatial profile")
	psiIm_k332		= NewExcitation("psiIm_k332", "", "Eigenmode spatial profile")
	psiIm_k333		= NewExcitation("psiIm_k333", "", "Eigenmode spatial profile")
	psiIm_k334		= NewExcitation("psiIm_k334", "", "Eigenmode spatial profile")
	psiIm_k335		= NewExcitation("psiIm_k335", "", "Eigenmode spatial profile")
	psiIm_k336		= NewExcitation("psiIm_k336", "", "Eigenmode spatial profile")
	psiIm_k337		= NewExcitation("psiIm_k337", "", "Eigenmode spatial profile")
	psiIm_k338		= NewExcitation("psiIm_k338", "", "Eigenmode spatial profile")
	psiIm_k339		= NewExcitation("psiIm_k339", "", "Eigenmode spatial profile")
	psiIm_k340		= NewExcitation("psiIm_k340", "", "Eigenmode spatial profile")
	psiIm_k341		= NewExcitation("psiIm_k341", "", "Eigenmode spatial profile")
	psiIm_k342		= NewExcitation("psiIm_k342", "", "Eigenmode spatial profile")
	psiIm_k343		= NewExcitation("psiIm_k343", "", "Eigenmode spatial profile")
	psiIm_k344		= NewExcitation("psiIm_k344", "", "Eigenmode spatial profile")
	psiIm_k345		= NewExcitation("psiIm_k345", "", "Eigenmode spatial profile")
	psiIm_k346		= NewExcitation("psiIm_k346", "", "Eigenmode spatial profile")
	psiIm_k347		= NewExcitation("psiIm_k347", "", "Eigenmode spatial profile")
	psiIm_k348		= NewExcitation("psiIm_k348", "", "Eigenmode spatial profile")
	psiIm_k349		= NewExcitation("psiIm_k349", "", "Eigenmode spatial profile")
	psiIm_k350		= NewExcitation("psiIm_k350", "", "Eigenmode spatial profile")
	psiIm_k351		= NewExcitation("psiIm_k351", "", "Eigenmode spatial profile")
	psiIm_k352		= NewExcitation("psiIm_k352", "", "Eigenmode spatial profile")
	psiIm_k353		= NewExcitation("psiIm_k353", "", "Eigenmode spatial profile")
	psiIm_k354		= NewExcitation("psiIm_k354", "", "Eigenmode spatial profile")
	psiIm_k355		= NewExcitation("psiIm_k355", "", "Eigenmode spatial profile")
	psiIm_k356		= NewExcitation("psiIm_k356", "", "Eigenmode spatial profile")
	psiIm_k357		= NewExcitation("psiIm_k357", "", "Eigenmode spatial profile")
	psiIm_k358		= NewExcitation("psiIm_k358", "", "Eigenmode spatial profile")
	psiIm_k359		= NewExcitation("psiIm_k359", "", "Eigenmode spatial profile")
	psiIm_k360		= NewExcitation("psiIm_k360", "", "Eigenmode spatial profile")
	psiIm_k361		= NewExcitation("psiIm_k361", "", "Eigenmode spatial profile")
	psiIm_k362		= NewExcitation("psiIm_k362", "", "Eigenmode spatial profile")
	psiIm_k363		= NewExcitation("psiIm_k363", "", "Eigenmode spatial profile")
	psiIm_k364		= NewExcitation("psiIm_k364", "", "Eigenmode spatial profile")
	psiIm_k365		= NewExcitation("psiIm_k365", "", "Eigenmode spatial profile")
	psiIm_k366		= NewExcitation("psiIm_k366", "", "Eigenmode spatial profile")
	psiIm_k367		= NewExcitation("psiIm_k367", "", "Eigenmode spatial profile")
	psiIm_k368		= NewExcitation("psiIm_k368", "", "Eigenmode spatial profile")
	psiIm_k369		= NewExcitation("psiIm_k369", "", "Eigenmode spatial profile")
	psiIm_k370		= NewExcitation("psiIm_k370", "", "Eigenmode spatial profile")
	psiIm_k371		= NewExcitation("psiIm_k371", "", "Eigenmode spatial profile")
	psiIm_k372		= NewExcitation("psiIm_k372", "", "Eigenmode spatial profile")
	psiIm_k373		= NewExcitation("psiIm_k373", "", "Eigenmode spatial profile")
	psiIm_k374		= NewExcitation("psiIm_k374", "", "Eigenmode spatial profile")
	psiIm_k375		= NewExcitation("psiIm_k375", "", "Eigenmode spatial profile")
	psiIm_k376		= NewExcitation("psiIm_k376", "", "Eigenmode spatial profile")
	psiIm_k377		= NewExcitation("psiIm_k377", "", "Eigenmode spatial profile")
	psiIm_k378		= NewExcitation("psiIm_k378", "", "Eigenmode spatial profile")
	psiIm_k379		= NewExcitation("psiIm_k379", "", "Eigenmode spatial profile")
	psiIm_k380		= NewExcitation("psiIm_k380", "", "Eigenmode spatial profile")
	psiIm_k381		= NewExcitation("psiIm_k381", "", "Eigenmode spatial profile")
	psiIm_k382		= NewExcitation("psiIm_k382", "", "Eigenmode spatial profile")
	psiIm_k383		= NewExcitation("psiIm_k383", "", "Eigenmode spatial profile")
	psiIm_k384		= NewExcitation("psiIm_k384", "", "Eigenmode spatial profile")
	psiIm_k385		= NewExcitation("psiIm_k385", "", "Eigenmode spatial profile")
	psiIm_k386		= NewExcitation("psiIm_k386", "", "Eigenmode spatial profile")
	psiIm_k387		= NewExcitation("psiIm_k387", "", "Eigenmode spatial profile")
	psiIm_k388		= NewExcitation("psiIm_k388", "", "Eigenmode spatial profile")
	psiIm_k389		= NewExcitation("psiIm_k389", "", "Eigenmode spatial profile")
	psiIm_k390		= NewExcitation("psiIm_k390", "", "Eigenmode spatial profile")
	psiIm_k391		= NewExcitation("psiIm_k391", "", "Eigenmode spatial profile")
	psiIm_k392		= NewExcitation("psiIm_k392", "", "Eigenmode spatial profile")
	psiIm_k393		= NewExcitation("psiIm_k393", "", "Eigenmode spatial profile")
	psiIm_k394		= NewExcitation("psiIm_k394", "", "Eigenmode spatial profile")
	psiIm_k395		= NewExcitation("psiIm_k395", "", "Eigenmode spatial profile")
	psiIm_k396		= NewExcitation("psiIm_k396", "", "Eigenmode spatial profile")
	psiIm_k397		= NewExcitation("psiIm_k397", "", "Eigenmode spatial profile")
	psiIm_k398		= NewExcitation("psiIm_k398", "", "Eigenmode spatial profile")
	psiIm_k399		= NewExcitation("psiIm_k399", "", "Eigenmode spatial profile")
	psiIm_k400		= NewExcitation("psiIm_k400", "", "Eigenmode spatial profile")
	psiIm_k401		= NewExcitation("psiIm_k401", "", "Eigenmode spatial profile")
	psiIm_k402		= NewExcitation("psiIm_k402", "", "Eigenmode spatial profile")
	psiIm_k403		= NewExcitation("psiIm_k403", "", "Eigenmode spatial profile")
	psiIm_k404		= NewExcitation("psiIm_k404", "", "Eigenmode spatial profile")
	psiIm_k405		= NewExcitation("psiIm_k405", "", "Eigenmode spatial profile")
	psiIm_k406		= NewExcitation("psiIm_k406", "", "Eigenmode spatial profile")
	psiIm_k407		= NewExcitation("psiIm_k407", "", "Eigenmode spatial profile")
	psiIm_k408		= NewExcitation("psiIm_k408", "", "Eigenmode spatial profile")
	psiIm_k409		= NewExcitation("psiIm_k409", "", "Eigenmode spatial profile")
	psiIm_k410		= NewExcitation("psiIm_k410", "", "Eigenmode spatial profile")
	psiIm_k411		= NewExcitation("psiIm_k411", "", "Eigenmode spatial profile")
	psiIm_k412		= NewExcitation("psiIm_k412", "", "Eigenmode spatial profile")
	psiIm_k413		= NewExcitation("psiIm_k413", "", "Eigenmode spatial profile")
	psiIm_k414		= NewExcitation("psiIm_k414", "", "Eigenmode spatial profile")
	psiIm_k415		= NewExcitation("psiIm_k415", "", "Eigenmode spatial profile")
	psiIm_k416		= NewExcitation("psiIm_k416", "", "Eigenmode spatial profile")
	psiIm_k417		= NewExcitation("psiIm_k417", "", "Eigenmode spatial profile")
	psiIm_k418		= NewExcitation("psiIm_k418", "", "Eigenmode spatial profile")
	psiIm_k419		= NewExcitation("psiIm_k419", "", "Eigenmode spatial profile")
	psiIm_k420		= NewExcitation("psiIm_k420", "", "Eigenmode spatial profile")
	psiIm_k421		= NewExcitation("psiIm_k421", "", "Eigenmode spatial profile")
	psiIm_k422		= NewExcitation("psiIm_k422", "", "Eigenmode spatial profile")
	psiIm_k423		= NewExcitation("psiIm_k423", "", "Eigenmode spatial profile")
	psiIm_k424		= NewExcitation("psiIm_k424", "", "Eigenmode spatial profile")
	psiIm_k425		= NewExcitation("psiIm_k425", "", "Eigenmode spatial profile")
	psiIm_k426		= NewExcitation("psiIm_k426", "", "Eigenmode spatial profile")
	psiIm_k427		= NewExcitation("psiIm_k427", "", "Eigenmode spatial profile")
	psiIm_k428		= NewExcitation("psiIm_k428", "", "Eigenmode spatial profile")
	psiIm_k429		= NewExcitation("psiIm_k429", "", "Eigenmode spatial profile")
	psiIm_k430		= NewExcitation("psiIm_k430", "", "Eigenmode spatial profile")
	psiIm_k431		= NewExcitation("psiIm_k431", "", "Eigenmode spatial profile")
	psiIm_k432		= NewExcitation("psiIm_k432", "", "Eigenmode spatial profile")
	psiIm_k433		= NewExcitation("psiIm_k433", "", "Eigenmode spatial profile")
	psiIm_k434		= NewExcitation("psiIm_k434", "", "Eigenmode spatial profile")
	psiIm_k435		= NewExcitation("psiIm_k435", "", "Eigenmode spatial profile")
	psiIm_k436		= NewExcitation("psiIm_k436", "", "Eigenmode spatial profile")
	psiIm_k437		= NewExcitation("psiIm_k437", "", "Eigenmode spatial profile")
	psiIm_k438		= NewExcitation("psiIm_k438", "", "Eigenmode spatial profile")
	psiIm_k439		= NewExcitation("psiIm_k439", "", "Eigenmode spatial profile")
	psiIm_k440		= NewExcitation("psiIm_k440", "", "Eigenmode spatial profile")
	psiIm_k441		= NewExcitation("psiIm_k441", "", "Eigenmode spatial profile")
	psiIm_k442		= NewExcitation("psiIm_k442", "", "Eigenmode spatial profile")
	psiIm_k443		= NewExcitation("psiIm_k443", "", "Eigenmode spatial profile")
	psiIm_k444		= NewExcitation("psiIm_k444", "", "Eigenmode spatial profile")
	psiIm_k445		= NewExcitation("psiIm_k445", "", "Eigenmode spatial profile")
	psiIm_k446		= NewExcitation("psiIm_k446", "", "Eigenmode spatial profile")
	psiIm_k447		= NewExcitation("psiIm_k447", "", "Eigenmode spatial profile")
	psiIm_k448		= NewExcitation("psiIm_k448", "", "Eigenmode spatial profile")
	psiIm_k449		= NewExcitation("psiIm_k449", "", "Eigenmode spatial profile")
	psiIm_k450		= NewExcitation("psiIm_k450", "", "Eigenmode spatial profile")
	psiIm_k451		= NewExcitation("psiIm_k451", "", "Eigenmode spatial profile")
	psiIm_k452		= NewExcitation("psiIm_k452", "", "Eigenmode spatial profile")
	psiIm_k453		= NewExcitation("psiIm_k453", "", "Eigenmode spatial profile")
	psiIm_k454		= NewExcitation("psiIm_k454", "", "Eigenmode spatial profile")
	psiIm_k455		= NewExcitation("psiIm_k455", "", "Eigenmode spatial profile")
	psiIm_k456		= NewExcitation("psiIm_k456", "", "Eigenmode spatial profile")
	psiIm_k457		= NewExcitation("psiIm_k457", "", "Eigenmode spatial profile")
	psiIm_k458		= NewExcitation("psiIm_k458", "", "Eigenmode spatial profile")
	psiIm_k459		= NewExcitation("psiIm_k459", "", "Eigenmode spatial profile")
	psiIm_k460		= NewExcitation("psiIm_k460", "", "Eigenmode spatial profile")
	psiIm_k461		= NewExcitation("psiIm_k461", "", "Eigenmode spatial profile")
	psiIm_k462		= NewExcitation("psiIm_k462", "", "Eigenmode spatial profile")
	psiIm_k463		= NewExcitation("psiIm_k463", "", "Eigenmode spatial profile")
	psiIm_k464		= NewExcitation("psiIm_k464", "", "Eigenmode spatial profile")
	psiIm_k465		= NewExcitation("psiIm_k465", "", "Eigenmode spatial profile")
	psiIm_k466		= NewExcitation("psiIm_k466", "", "Eigenmode spatial profile")
	psiIm_k467		= NewExcitation("psiIm_k467", "", "Eigenmode spatial profile")
	psiIm_k468		= NewExcitation("psiIm_k468", "", "Eigenmode spatial profile")
	psiIm_k469		= NewExcitation("psiIm_k469", "", "Eigenmode spatial profile")
	psiIm_k470		= NewExcitation("psiIm_k470", "", "Eigenmode spatial profile")
	psiIm_k471		= NewExcitation("psiIm_k471", "", "Eigenmode spatial profile")
	psiIm_k472		= NewExcitation("psiIm_k472", "", "Eigenmode spatial profile")
	psiIm_k473		= NewExcitation("psiIm_k473", "", "Eigenmode spatial profile")
	psiIm_k474		= NewExcitation("psiIm_k474", "", "Eigenmode spatial profile")
	psiIm_k475		= NewExcitation("psiIm_k475", "", "Eigenmode spatial profile")
	psiIm_k476		= NewExcitation("psiIm_k476", "", "Eigenmode spatial profile")
	psiIm_k477		= NewExcitation("psiIm_k477", "", "Eigenmode spatial profile")
	psiIm_k478		= NewExcitation("psiIm_k478", "", "Eigenmode spatial profile")
	psiIm_k479		= NewExcitation("psiIm_k479", "", "Eigenmode spatial profile")
	psiIm_k480		= NewExcitation("psiIm_k480", "", "Eigenmode spatial profile")
	psiIm_k481		= NewExcitation("psiIm_k481", "", "Eigenmode spatial profile")
	psiIm_k482		= NewExcitation("psiIm_k482", "", "Eigenmode spatial profile")
	psiIm_k483		= NewExcitation("psiIm_k483", "", "Eigenmode spatial profile")
	psiIm_k484		= NewExcitation("psiIm_k484", "", "Eigenmode spatial profile")
	psiIm_k485		= NewExcitation("psiIm_k485", "", "Eigenmode spatial profile")
	psiIm_k486		= NewExcitation("psiIm_k486", "", "Eigenmode spatial profile")
	psiIm_k487		= NewExcitation("psiIm_k487", "", "Eigenmode spatial profile")
	psiIm_k488		= NewExcitation("psiIm_k488", "", "Eigenmode spatial profile")
	psiIm_k489		= NewExcitation("psiIm_k489", "", "Eigenmode spatial profile")
	psiIm_k490		= NewExcitation("psiIm_k490", "", "Eigenmode spatial profile")
	psiIm_k491		= NewExcitation("psiIm_k491", "", "Eigenmode spatial profile")
	psiIm_k492		= NewExcitation("psiIm_k492", "", "Eigenmode spatial profile")
	psiIm_k493		= NewExcitation("psiIm_k493", "", "Eigenmode spatial profile")
	psiIm_k494		= NewExcitation("psiIm_k494", "", "Eigenmode spatial profile")
	psiIm_k495		= NewExcitation("psiIm_k495", "", "Eigenmode spatial profile")
	psiIm_k496		= NewExcitation("psiIm_k496", "", "Eigenmode spatial profile")
	psiIm_k497		= NewExcitation("psiIm_k497", "", "Eigenmode spatial profile")
	psiIm_k498		= NewExcitation("psiIm_k498", "", "Eigenmode spatial profile")
	psiIm_k499		= NewExcitation("psiIm_k499", "", "Eigenmode spatial profile")
	
	b_k000		= NewVectorValue("b_k000", "", "m projection onto psi(Re,Im)_k000", GetModeAmplitudeReImk000 )
	b_k001		= NewVectorValue("b_k001", "", "m projection onto psi(Re,Im)_k001", GetModeAmplitudeReImk001 )
	b_k002		= NewVectorValue("b_k002", "", "m projection onto psi(Re,Im)_k002", GetModeAmplitudeReImk002 )
	b_k003		= NewVectorValue("b_k003", "", "m projection onto psi(Re,Im)_k003", GetModeAmplitudeReImk003 )
	b_k004		= NewVectorValue("b_k004", "", "m projection onto psi(Re,Im)_k004", GetModeAmplitudeReImk004 )
	b_k005		= NewVectorValue("b_k005", "", "m projection onto psi(Re,Im)_k005", GetModeAmplitudeReImk005 )
	b_k006		= NewVectorValue("b_k006", "", "m projection onto psi(Re,Im)_k006", GetModeAmplitudeReImk006 )
	b_k007		= NewVectorValue("b_k007", "", "m projection onto psi(Re,Im)_k007", GetModeAmplitudeReImk007 )
	b_k008		= NewVectorValue("b_k008", "", "m projection onto psi(Re,Im)_k008", GetModeAmplitudeReImk008 )
	b_k009		= NewVectorValue("b_k009", "", "m projection onto psi(Re,Im)_k009", GetModeAmplitudeReImk009 )
	b_k010		= NewVectorValue("b_k010", "", "m projection onto psi(Re,Im)_k010", GetModeAmplitudeReImk010 )
	b_k011		= NewVectorValue("b_k011", "", "m projection onto psi(Re,Im)_k011", GetModeAmplitudeReImk011 )
	b_k012		= NewVectorValue("b_k012", "", "m projection onto psi(Re,Im)_k012", GetModeAmplitudeReImk012 )
	b_k013		= NewVectorValue("b_k013", "", "m projection onto psi(Re,Im)_k013", GetModeAmplitudeReImk013 )
	b_k014		= NewVectorValue("b_k014", "", "m projection onto psi(Re,Im)_k014", GetModeAmplitudeReImk014 )
	b_k015		= NewVectorValue("b_k015", "", "m projection onto psi(Re,Im)_k015", GetModeAmplitudeReImk015 )
	b_k016		= NewVectorValue("b_k016", "", "m projection onto psi(Re,Im)_k016", GetModeAmplitudeReImk016 )
	b_k017		= NewVectorValue("b_k017", "", "m projection onto psi(Re,Im)_k017", GetModeAmplitudeReImk017 )
	b_k018		= NewVectorValue("b_k018", "", "m projection onto psi(Re,Im)_k018", GetModeAmplitudeReImk018 )
	b_k019		= NewVectorValue("b_k019", "", "m projection onto psi(Re,Im)_k019", GetModeAmplitudeReImk019 )
	b_k020		= NewVectorValue("b_k020", "", "m projection onto psi(Re,Im)_k020", GetModeAmplitudeReImk020 )
	b_k021		= NewVectorValue("b_k021", "", "m projection onto psi(Re,Im)_k021", GetModeAmplitudeReImk021 )
	b_k022		= NewVectorValue("b_k022", "", "m projection onto psi(Re,Im)_k022", GetModeAmplitudeReImk022 )
	b_k023		= NewVectorValue("b_k023", "", "m projection onto psi(Re,Im)_k023", GetModeAmplitudeReImk023 )
	b_k024		= NewVectorValue("b_k024", "", "m projection onto psi(Re,Im)_k024", GetModeAmplitudeReImk024 )
	b_k025		= NewVectorValue("b_k025", "", "m projection onto psi(Re,Im)_k025", GetModeAmplitudeReImk025 )
	b_k026		= NewVectorValue("b_k026", "", "m projection onto psi(Re,Im)_k026", GetModeAmplitudeReImk026 )
	b_k027		= NewVectorValue("b_k027", "", "m projection onto psi(Re,Im)_k027", GetModeAmplitudeReImk027 )
	b_k028		= NewVectorValue("b_k028", "", "m projection onto psi(Re,Im)_k028", GetModeAmplitudeReImk028 )
	b_k029		= NewVectorValue("b_k029", "", "m projection onto psi(Re,Im)_k029", GetModeAmplitudeReImk029 )
	b_k030		= NewVectorValue("b_k030", "", "m projection onto psi(Re,Im)_k030", GetModeAmplitudeReImk030 )
	b_k031		= NewVectorValue("b_k031", "", "m projection onto psi(Re,Im)_k031", GetModeAmplitudeReImk031 )
	b_k032		= NewVectorValue("b_k032", "", "m projection onto psi(Re,Im)_k032", GetModeAmplitudeReImk032 )
	b_k033		= NewVectorValue("b_k033", "", "m projection onto psi(Re,Im)_k033", GetModeAmplitudeReImk033 )
	b_k034		= NewVectorValue("b_k034", "", "m projection onto psi(Re,Im)_k034", GetModeAmplitudeReImk034 )
	b_k035		= NewVectorValue("b_k035", "", "m projection onto psi(Re,Im)_k035", GetModeAmplitudeReImk035 )
	b_k036		= NewVectorValue("b_k036", "", "m projection onto psi(Re,Im)_k036", GetModeAmplitudeReImk036 )
	b_k037		= NewVectorValue("b_k037", "", "m projection onto psi(Re,Im)_k037", GetModeAmplitudeReImk037 )
	b_k038		= NewVectorValue("b_k038", "", "m projection onto psi(Re,Im)_k038", GetModeAmplitudeReImk038 )
	b_k039		= NewVectorValue("b_k039", "", "m projection onto psi(Re,Im)_k039", GetModeAmplitudeReImk039 )
	b_k040		= NewVectorValue("b_k040", "", "m projection onto psi(Re,Im)_k040", GetModeAmplitudeReImk040 )
	b_k041		= NewVectorValue("b_k041", "", "m projection onto psi(Re,Im)_k041", GetModeAmplitudeReImk041 )
	b_k042		= NewVectorValue("b_k042", "", "m projection onto psi(Re,Im)_k042", GetModeAmplitudeReImk042 )
	b_k043		= NewVectorValue("b_k043", "", "m projection onto psi(Re,Im)_k043", GetModeAmplitudeReImk043 )
	b_k044		= NewVectorValue("b_k044", "", "m projection onto psi(Re,Im)_k044", GetModeAmplitudeReImk044 )
	b_k045		= NewVectorValue("b_k045", "", "m projection onto psi(Re,Im)_k045", GetModeAmplitudeReImk045 )
	b_k046		= NewVectorValue("b_k046", "", "m projection onto psi(Re,Im)_k046", GetModeAmplitudeReImk046 )
	b_k047		= NewVectorValue("b_k047", "", "m projection onto psi(Re,Im)_k047", GetModeAmplitudeReImk047 )
	b_k048		= NewVectorValue("b_k048", "", "m projection onto psi(Re,Im)_k048", GetModeAmplitudeReImk048 )
	b_k049		= NewVectorValue("b_k049", "", "m projection onto psi(Re,Im)_k049", GetModeAmplitudeReImk049 )
	b_k050		= NewVectorValue("b_k050", "", "m projection onto psi(Re,Im)_k050", GetModeAmplitudeReImk050 )
	b_k051		= NewVectorValue("b_k051", "", "m projection onto psi(Re,Im)_k051", GetModeAmplitudeReImk051 )
	b_k052		= NewVectorValue("b_k052", "", "m projection onto psi(Re,Im)_k052", GetModeAmplitudeReImk052 )
	b_k053		= NewVectorValue("b_k053", "", "m projection onto psi(Re,Im)_k053", GetModeAmplitudeReImk053 )
	b_k054		= NewVectorValue("b_k054", "", "m projection onto psi(Re,Im)_k054", GetModeAmplitudeReImk054 )
	b_k055		= NewVectorValue("b_k055", "", "m projection onto psi(Re,Im)_k055", GetModeAmplitudeReImk055 )
	b_k056		= NewVectorValue("b_k056", "", "m projection onto psi(Re,Im)_k056", GetModeAmplitudeReImk056 )
	b_k057		= NewVectorValue("b_k057", "", "m projection onto psi(Re,Im)_k057", GetModeAmplitudeReImk057 )
	b_k058		= NewVectorValue("b_k058", "", "m projection onto psi(Re,Im)_k058", GetModeAmplitudeReImk058 )
	b_k059		= NewVectorValue("b_k059", "", "m projection onto psi(Re,Im)_k059", GetModeAmplitudeReImk059 )
	b_k060		= NewVectorValue("b_k060", "", "m projection onto psi(Re,Im)_k060", GetModeAmplitudeReImk060 )
	b_k061		= NewVectorValue("b_k061", "", "m projection onto psi(Re,Im)_k061", GetModeAmplitudeReImk061 )
	b_k062		= NewVectorValue("b_k062", "", "m projection onto psi(Re,Im)_k062", GetModeAmplitudeReImk062 )
	b_k063		= NewVectorValue("b_k063", "", "m projection onto psi(Re,Im)_k063", GetModeAmplitudeReImk063 )
	b_k064		= NewVectorValue("b_k064", "", "m projection onto psi(Re,Im)_k064", GetModeAmplitudeReImk064 )
	b_k065		= NewVectorValue("b_k065", "", "m projection onto psi(Re,Im)_k065", GetModeAmplitudeReImk065 )
	b_k066		= NewVectorValue("b_k066", "", "m projection onto psi(Re,Im)_k066", GetModeAmplitudeReImk066 )
	b_k067		= NewVectorValue("b_k067", "", "m projection onto psi(Re,Im)_k067", GetModeAmplitudeReImk067 )
	b_k068		= NewVectorValue("b_k068", "", "m projection onto psi(Re,Im)_k068", GetModeAmplitudeReImk068 )
	b_k069		= NewVectorValue("b_k069", "", "m projection onto psi(Re,Im)_k069", GetModeAmplitudeReImk069 )
	b_k070		= NewVectorValue("b_k070", "", "m projection onto psi(Re,Im)_k070", GetModeAmplitudeReImk070 )
	b_k071		= NewVectorValue("b_k071", "", "m projection onto psi(Re,Im)_k071", GetModeAmplitudeReImk071 )
	b_k072		= NewVectorValue("b_k072", "", "m projection onto psi(Re,Im)_k072", GetModeAmplitudeReImk072 )
	b_k073		= NewVectorValue("b_k073", "", "m projection onto psi(Re,Im)_k073", GetModeAmplitudeReImk073 )
	b_k074		= NewVectorValue("b_k074", "", "m projection onto psi(Re,Im)_k074", GetModeAmplitudeReImk074 )
	b_k075		= NewVectorValue("b_k075", "", "m projection onto psi(Re,Im)_k075", GetModeAmplitudeReImk075 )
	b_k076		= NewVectorValue("b_k076", "", "m projection onto psi(Re,Im)_k076", GetModeAmplitudeReImk076 )
	b_k077		= NewVectorValue("b_k077", "", "m projection onto psi(Re,Im)_k077", GetModeAmplitudeReImk077 )
	b_k078		= NewVectorValue("b_k078", "", "m projection onto psi(Re,Im)_k078", GetModeAmplitudeReImk078 )
	b_k079		= NewVectorValue("b_k079", "", "m projection onto psi(Re,Im)_k079", GetModeAmplitudeReImk079 )
	b_k080		= NewVectorValue("b_k080", "", "m projection onto psi(Re,Im)_k080", GetModeAmplitudeReImk080 )
	b_k081		= NewVectorValue("b_k081", "", "m projection onto psi(Re,Im)_k081", GetModeAmplitudeReImk081 )
	b_k082		= NewVectorValue("b_k082", "", "m projection onto psi(Re,Im)_k082", GetModeAmplitudeReImk082 )
	b_k083		= NewVectorValue("b_k083", "", "m projection onto psi(Re,Im)_k083", GetModeAmplitudeReImk083 )
	b_k084		= NewVectorValue("b_k084", "", "m projection onto psi(Re,Im)_k084", GetModeAmplitudeReImk084 )
	b_k085		= NewVectorValue("b_k085", "", "m projection onto psi(Re,Im)_k085", GetModeAmplitudeReImk085 )
	b_k086		= NewVectorValue("b_k086", "", "m projection onto psi(Re,Im)_k086", GetModeAmplitudeReImk086 )
	b_k087		= NewVectorValue("b_k087", "", "m projection onto psi(Re,Im)_k087", GetModeAmplitudeReImk087 )
	b_k088		= NewVectorValue("b_k088", "", "m projection onto psi(Re,Im)_k088", GetModeAmplitudeReImk088 )
	b_k089		= NewVectorValue("b_k089", "", "m projection onto psi(Re,Im)_k089", GetModeAmplitudeReImk089 )
	b_k090		= NewVectorValue("b_k090", "", "m projection onto psi(Re,Im)_k090", GetModeAmplitudeReImk090 )
	b_k091		= NewVectorValue("b_k091", "", "m projection onto psi(Re,Im)_k091", GetModeAmplitudeReImk091 )
	b_k092		= NewVectorValue("b_k092", "", "m projection onto psi(Re,Im)_k092", GetModeAmplitudeReImk092 )
	b_k093		= NewVectorValue("b_k093", "", "m projection onto psi(Re,Im)_k093", GetModeAmplitudeReImk093 )
	b_k094		= NewVectorValue("b_k094", "", "m projection onto psi(Re,Im)_k094", GetModeAmplitudeReImk094 )
	b_k095		= NewVectorValue("b_k095", "", "m projection onto psi(Re,Im)_k095", GetModeAmplitudeReImk095 )
	b_k096		= NewVectorValue("b_k096", "", "m projection onto psi(Re,Im)_k096", GetModeAmplitudeReImk096 )
	b_k097		= NewVectorValue("b_k097", "", "m projection onto psi(Re,Im)_k097", GetModeAmplitudeReImk097 )
	b_k098		= NewVectorValue("b_k098", "", "m projection onto psi(Re,Im)_k098", GetModeAmplitudeReImk098 )
	b_k099		= NewVectorValue("b_k099", "", "m projection onto psi(Re,Im)_k099", GetModeAmplitudeReImk099 )
	b_k100		= NewVectorValue("b_k100", "", "m projection onto psi(Re,Im)_k100", GetModeAmplitudeReImk100 )
	b_k101		= NewVectorValue("b_k101", "", "m projection onto psi(Re,Im)_k101", GetModeAmplitudeReImk101 )
	b_k102		= NewVectorValue("b_k102", "", "m projection onto psi(Re,Im)_k102", GetModeAmplitudeReImk102 )
	b_k103		= NewVectorValue("b_k103", "", "m projection onto psi(Re,Im)_k103", GetModeAmplitudeReImk103 )
	b_k104		= NewVectorValue("b_k104", "", "m projection onto psi(Re,Im)_k104", GetModeAmplitudeReImk104 )
	b_k105		= NewVectorValue("b_k105", "", "m projection onto psi(Re,Im)_k105", GetModeAmplitudeReImk105 )
	b_k106		= NewVectorValue("b_k106", "", "m projection onto psi(Re,Im)_k106", GetModeAmplitudeReImk106 )
	b_k107		= NewVectorValue("b_k107", "", "m projection onto psi(Re,Im)_k107", GetModeAmplitudeReImk107 )
	b_k108		= NewVectorValue("b_k108", "", "m projection onto psi(Re,Im)_k108", GetModeAmplitudeReImk108 )
	b_k109		= NewVectorValue("b_k109", "", "m projection onto psi(Re,Im)_k109", GetModeAmplitudeReImk109 )
	b_k110		= NewVectorValue("b_k110", "", "m projection onto psi(Re,Im)_k110", GetModeAmplitudeReImk110 )
	b_k111		= NewVectorValue("b_k111", "", "m projection onto psi(Re,Im)_k111", GetModeAmplitudeReImk111 )
	b_k112		= NewVectorValue("b_k112", "", "m projection onto psi(Re,Im)_k112", GetModeAmplitudeReImk112 )
	b_k113		= NewVectorValue("b_k113", "", "m projection onto psi(Re,Im)_k113", GetModeAmplitudeReImk113 )
	b_k114		= NewVectorValue("b_k114", "", "m projection onto psi(Re,Im)_k114", GetModeAmplitudeReImk114 )
	b_k115		= NewVectorValue("b_k115", "", "m projection onto psi(Re,Im)_k115", GetModeAmplitudeReImk115 )
	b_k116		= NewVectorValue("b_k116", "", "m projection onto psi(Re,Im)_k116", GetModeAmplitudeReImk116 )
	b_k117		= NewVectorValue("b_k117", "", "m projection onto psi(Re,Im)_k117", GetModeAmplitudeReImk117 )
	b_k118		= NewVectorValue("b_k118", "", "m projection onto psi(Re,Im)_k118", GetModeAmplitudeReImk118 )
	b_k119		= NewVectorValue("b_k119", "", "m projection onto psi(Re,Im)_k119", GetModeAmplitudeReImk119 )
	b_k120		= NewVectorValue("b_k120", "", "m projection onto psi(Re,Im)_k120", GetModeAmplitudeReImk120 )
	b_k121		= NewVectorValue("b_k121", "", "m projection onto psi(Re,Im)_k121", GetModeAmplitudeReImk121 )
	b_k122		= NewVectorValue("b_k122", "", "m projection onto psi(Re,Im)_k122", GetModeAmplitudeReImk122 )
	b_k123		= NewVectorValue("b_k123", "", "m projection onto psi(Re,Im)_k123", GetModeAmplitudeReImk123 )
	b_k124		= NewVectorValue("b_k124", "", "m projection onto psi(Re,Im)_k124", GetModeAmplitudeReImk124 )
	b_k125		= NewVectorValue("b_k125", "", "m projection onto psi(Re,Im)_k125", GetModeAmplitudeReImk125 )
	b_k126		= NewVectorValue("b_k126", "", "m projection onto psi(Re,Im)_k126", GetModeAmplitudeReImk126 )
	b_k127		= NewVectorValue("b_k127", "", "m projection onto psi(Re,Im)_k127", GetModeAmplitudeReImk127 )
	b_k128		= NewVectorValue("b_k128", "", "m projection onto psi(Re,Im)_k128", GetModeAmplitudeReImk128 )
	b_k129		= NewVectorValue("b_k129", "", "m projection onto psi(Re,Im)_k129", GetModeAmplitudeReImk129 )
	b_k130		= NewVectorValue("b_k130", "", "m projection onto psi(Re,Im)_k130", GetModeAmplitudeReImk130 )
	b_k131		= NewVectorValue("b_k131", "", "m projection onto psi(Re,Im)_k131", GetModeAmplitudeReImk131 )
	b_k132		= NewVectorValue("b_k132", "", "m projection onto psi(Re,Im)_k132", GetModeAmplitudeReImk132 )
	b_k133		= NewVectorValue("b_k133", "", "m projection onto psi(Re,Im)_k133", GetModeAmplitudeReImk133 )
	b_k134		= NewVectorValue("b_k134", "", "m projection onto psi(Re,Im)_k134", GetModeAmplitudeReImk134 )
	b_k135		= NewVectorValue("b_k135", "", "m projection onto psi(Re,Im)_k135", GetModeAmplitudeReImk135 )
	b_k136		= NewVectorValue("b_k136", "", "m projection onto psi(Re,Im)_k136", GetModeAmplitudeReImk136 )
	b_k137		= NewVectorValue("b_k137", "", "m projection onto psi(Re,Im)_k137", GetModeAmplitudeReImk137 )
	b_k138		= NewVectorValue("b_k138", "", "m projection onto psi(Re,Im)_k138", GetModeAmplitudeReImk138 )
	b_k139		= NewVectorValue("b_k139", "", "m projection onto psi(Re,Im)_k139", GetModeAmplitudeReImk139 )
	b_k140		= NewVectorValue("b_k140", "", "m projection onto psi(Re,Im)_k140", GetModeAmplitudeReImk140 )
	b_k141		= NewVectorValue("b_k141", "", "m projection onto psi(Re,Im)_k141", GetModeAmplitudeReImk141 )
	b_k142		= NewVectorValue("b_k142", "", "m projection onto psi(Re,Im)_k142", GetModeAmplitudeReImk142 )
	b_k143		= NewVectorValue("b_k143", "", "m projection onto psi(Re,Im)_k143", GetModeAmplitudeReImk143 )
	b_k144		= NewVectorValue("b_k144", "", "m projection onto psi(Re,Im)_k144", GetModeAmplitudeReImk144 )
	b_k145		= NewVectorValue("b_k145", "", "m projection onto psi(Re,Im)_k145", GetModeAmplitudeReImk145 )
	b_k146		= NewVectorValue("b_k146", "", "m projection onto psi(Re,Im)_k146", GetModeAmplitudeReImk146 )
	b_k147		= NewVectorValue("b_k147", "", "m projection onto psi(Re,Im)_k147", GetModeAmplitudeReImk147 )
	b_k148		= NewVectorValue("b_k148", "", "m projection onto psi(Re,Im)_k148", GetModeAmplitudeReImk148 )
	b_k149		= NewVectorValue("b_k149", "", "m projection onto psi(Re,Im)_k149", GetModeAmplitudeReImk149 )
	b_k150		= NewVectorValue("b_k150", "", "m projection onto psi(Re,Im)_k150", GetModeAmplitudeReImk150 )
	b_k151		= NewVectorValue("b_k151", "", "m projection onto psi(Re,Im)_k151", GetModeAmplitudeReImk151 )
	b_k152		= NewVectorValue("b_k152", "", "m projection onto psi(Re,Im)_k152", GetModeAmplitudeReImk152 )
	b_k153		= NewVectorValue("b_k153", "", "m projection onto psi(Re,Im)_k153", GetModeAmplitudeReImk153 )
	b_k154		= NewVectorValue("b_k154", "", "m projection onto psi(Re,Im)_k154", GetModeAmplitudeReImk154 )
	b_k155		= NewVectorValue("b_k155", "", "m projection onto psi(Re,Im)_k155", GetModeAmplitudeReImk155 )
	b_k156		= NewVectorValue("b_k156", "", "m projection onto psi(Re,Im)_k156", GetModeAmplitudeReImk156 )
	b_k157		= NewVectorValue("b_k157", "", "m projection onto psi(Re,Im)_k157", GetModeAmplitudeReImk157 )
	b_k158		= NewVectorValue("b_k158", "", "m projection onto psi(Re,Im)_k158", GetModeAmplitudeReImk158 )
	b_k159		= NewVectorValue("b_k159", "", "m projection onto psi(Re,Im)_k159", GetModeAmplitudeReImk159 )
	b_k160		= NewVectorValue("b_k160", "", "m projection onto psi(Re,Im)_k160", GetModeAmplitudeReImk160 )
	b_k161		= NewVectorValue("b_k161", "", "m projection onto psi(Re,Im)_k161", GetModeAmplitudeReImk161 )
	b_k162		= NewVectorValue("b_k162", "", "m projection onto psi(Re,Im)_k162", GetModeAmplitudeReImk162 )
	b_k163		= NewVectorValue("b_k163", "", "m projection onto psi(Re,Im)_k163", GetModeAmplitudeReImk163 )
	b_k164		= NewVectorValue("b_k164", "", "m projection onto psi(Re,Im)_k164", GetModeAmplitudeReImk164 )
	b_k165		= NewVectorValue("b_k165", "", "m projection onto psi(Re,Im)_k165", GetModeAmplitudeReImk165 )
	b_k166		= NewVectorValue("b_k166", "", "m projection onto psi(Re,Im)_k166", GetModeAmplitudeReImk166 )
	b_k167		= NewVectorValue("b_k167", "", "m projection onto psi(Re,Im)_k167", GetModeAmplitudeReImk167 )
	b_k168		= NewVectorValue("b_k168", "", "m projection onto psi(Re,Im)_k168", GetModeAmplitudeReImk168 )
	b_k169		= NewVectorValue("b_k169", "", "m projection onto psi(Re,Im)_k169", GetModeAmplitudeReImk169 )
	b_k170		= NewVectorValue("b_k170", "", "m projection onto psi(Re,Im)_k170", GetModeAmplitudeReImk170 )
	b_k171		= NewVectorValue("b_k171", "", "m projection onto psi(Re,Im)_k171", GetModeAmplitudeReImk171 )
	b_k172		= NewVectorValue("b_k172", "", "m projection onto psi(Re,Im)_k172", GetModeAmplitudeReImk172 )
	b_k173		= NewVectorValue("b_k173", "", "m projection onto psi(Re,Im)_k173", GetModeAmplitudeReImk173 )
	b_k174		= NewVectorValue("b_k174", "", "m projection onto psi(Re,Im)_k174", GetModeAmplitudeReImk174 )
	b_k175		= NewVectorValue("b_k175", "", "m projection onto psi(Re,Im)_k175", GetModeAmplitudeReImk175 )
	b_k176		= NewVectorValue("b_k176", "", "m projection onto psi(Re,Im)_k176", GetModeAmplitudeReImk176 )
	b_k177		= NewVectorValue("b_k177", "", "m projection onto psi(Re,Im)_k177", GetModeAmplitudeReImk177 )
	b_k178		= NewVectorValue("b_k178", "", "m projection onto psi(Re,Im)_k178", GetModeAmplitudeReImk178 )
	b_k179		= NewVectorValue("b_k179", "", "m projection onto psi(Re,Im)_k179", GetModeAmplitudeReImk179 )
	b_k180		= NewVectorValue("b_k180", "", "m projection onto psi(Re,Im)_k180", GetModeAmplitudeReImk180 )
	b_k181		= NewVectorValue("b_k181", "", "m projection onto psi(Re,Im)_k181", GetModeAmplitudeReImk181 )
	b_k182		= NewVectorValue("b_k182", "", "m projection onto psi(Re,Im)_k182", GetModeAmplitudeReImk182 )
	b_k183		= NewVectorValue("b_k183", "", "m projection onto psi(Re,Im)_k183", GetModeAmplitudeReImk183 )
	b_k184		= NewVectorValue("b_k184", "", "m projection onto psi(Re,Im)_k184", GetModeAmplitudeReImk184 )
	b_k185		= NewVectorValue("b_k185", "", "m projection onto psi(Re,Im)_k185", GetModeAmplitudeReImk185 )
	b_k186		= NewVectorValue("b_k186", "", "m projection onto psi(Re,Im)_k186", GetModeAmplitudeReImk186 )
	b_k187		= NewVectorValue("b_k187", "", "m projection onto psi(Re,Im)_k187", GetModeAmplitudeReImk187 )
	b_k188		= NewVectorValue("b_k188", "", "m projection onto psi(Re,Im)_k188", GetModeAmplitudeReImk188 )
	b_k189		= NewVectorValue("b_k189", "", "m projection onto psi(Re,Im)_k189", GetModeAmplitudeReImk189 )
	b_k190		= NewVectorValue("b_k190", "", "m projection onto psi(Re,Im)_k190", GetModeAmplitudeReImk190 )
	b_k191		= NewVectorValue("b_k191", "", "m projection onto psi(Re,Im)_k191", GetModeAmplitudeReImk191 )
	b_k192		= NewVectorValue("b_k192", "", "m projection onto psi(Re,Im)_k192", GetModeAmplitudeReImk192 )
	b_k193		= NewVectorValue("b_k193", "", "m projection onto psi(Re,Im)_k193", GetModeAmplitudeReImk193 )
	b_k194		= NewVectorValue("b_k194", "", "m projection onto psi(Re,Im)_k194", GetModeAmplitudeReImk194 )
	b_k195		= NewVectorValue("b_k195", "", "m projection onto psi(Re,Im)_k195", GetModeAmplitudeReImk195 )
	b_k196		= NewVectorValue("b_k196", "", "m projection onto psi(Re,Im)_k196", GetModeAmplitudeReImk196 )
	b_k197		= NewVectorValue("b_k197", "", "m projection onto psi(Re,Im)_k197", GetModeAmplitudeReImk197 )
	b_k198		= NewVectorValue("b_k198", "", "m projection onto psi(Re,Im)_k198", GetModeAmplitudeReImk198 )
	b_k199		= NewVectorValue("b_k199", "", "m projection onto psi(Re,Im)_k199", GetModeAmplitudeReImk199 )
	b_k200		= NewVectorValue("b_k200", "", "m projection onto psi(Re,Im)_k200", GetModeAmplitudeReImk200 )
	b_k201		= NewVectorValue("b_k201", "", "m projection onto psi(Re,Im)_k201", GetModeAmplitudeReImk201 )
	b_k202		= NewVectorValue("b_k202", "", "m projection onto psi(Re,Im)_k202", GetModeAmplitudeReImk202 )
	b_k203		= NewVectorValue("b_k203", "", "m projection onto psi(Re,Im)_k203", GetModeAmplitudeReImk203 )
	b_k204		= NewVectorValue("b_k204", "", "m projection onto psi(Re,Im)_k204", GetModeAmplitudeReImk204 )
	b_k205		= NewVectorValue("b_k205", "", "m projection onto psi(Re,Im)_k205", GetModeAmplitudeReImk205 )
	b_k206		= NewVectorValue("b_k206", "", "m projection onto psi(Re,Im)_k206", GetModeAmplitudeReImk206 )
	b_k207		= NewVectorValue("b_k207", "", "m projection onto psi(Re,Im)_k207", GetModeAmplitudeReImk207 )
	b_k208		= NewVectorValue("b_k208", "", "m projection onto psi(Re,Im)_k208", GetModeAmplitudeReImk208 )
	b_k209		= NewVectorValue("b_k209", "", "m projection onto psi(Re,Im)_k209", GetModeAmplitudeReImk209 )
	b_k210		= NewVectorValue("b_k210", "", "m projection onto psi(Re,Im)_k210", GetModeAmplitudeReImk210 )
	b_k211		= NewVectorValue("b_k211", "", "m projection onto psi(Re,Im)_k211", GetModeAmplitudeReImk211 )
	b_k212		= NewVectorValue("b_k212", "", "m projection onto psi(Re,Im)_k212", GetModeAmplitudeReImk212 )
	b_k213		= NewVectorValue("b_k213", "", "m projection onto psi(Re,Im)_k213", GetModeAmplitudeReImk213 )
	b_k214		= NewVectorValue("b_k214", "", "m projection onto psi(Re,Im)_k214", GetModeAmplitudeReImk214 )
	b_k215		= NewVectorValue("b_k215", "", "m projection onto psi(Re,Im)_k215", GetModeAmplitudeReImk215 )
	b_k216		= NewVectorValue("b_k216", "", "m projection onto psi(Re,Im)_k216", GetModeAmplitudeReImk216 )
	b_k217		= NewVectorValue("b_k217", "", "m projection onto psi(Re,Im)_k217", GetModeAmplitudeReImk217 )
	b_k218		= NewVectorValue("b_k218", "", "m projection onto psi(Re,Im)_k218", GetModeAmplitudeReImk218 )
	b_k219		= NewVectorValue("b_k219", "", "m projection onto psi(Re,Im)_k219", GetModeAmplitudeReImk219 )
	b_k220		= NewVectorValue("b_k220", "", "m projection onto psi(Re,Im)_k220", GetModeAmplitudeReImk220 )
	b_k221		= NewVectorValue("b_k221", "", "m projection onto psi(Re,Im)_k221", GetModeAmplitudeReImk221 )
	b_k222		= NewVectorValue("b_k222", "", "m projection onto psi(Re,Im)_k222", GetModeAmplitudeReImk222 )
	b_k223		= NewVectorValue("b_k223", "", "m projection onto psi(Re,Im)_k223", GetModeAmplitudeReImk223 )
	b_k224		= NewVectorValue("b_k224", "", "m projection onto psi(Re,Im)_k224", GetModeAmplitudeReImk224 )
	b_k225		= NewVectorValue("b_k225", "", "m projection onto psi(Re,Im)_k225", GetModeAmplitudeReImk225 )
	b_k226		= NewVectorValue("b_k226", "", "m projection onto psi(Re,Im)_k226", GetModeAmplitudeReImk226 )
	b_k227		= NewVectorValue("b_k227", "", "m projection onto psi(Re,Im)_k227", GetModeAmplitudeReImk227 )
	b_k228		= NewVectorValue("b_k228", "", "m projection onto psi(Re,Im)_k228", GetModeAmplitudeReImk228 )
	b_k229		= NewVectorValue("b_k229", "", "m projection onto psi(Re,Im)_k229", GetModeAmplitudeReImk229 )
	b_k230		= NewVectorValue("b_k230", "", "m projection onto psi(Re,Im)_k230", GetModeAmplitudeReImk230 )
	b_k231		= NewVectorValue("b_k231", "", "m projection onto psi(Re,Im)_k231", GetModeAmplitudeReImk231 )
	b_k232		= NewVectorValue("b_k232", "", "m projection onto psi(Re,Im)_k232", GetModeAmplitudeReImk232 )
	b_k233		= NewVectorValue("b_k233", "", "m projection onto psi(Re,Im)_k233", GetModeAmplitudeReImk233 )
	b_k234		= NewVectorValue("b_k234", "", "m projection onto psi(Re,Im)_k234", GetModeAmplitudeReImk234 )
	b_k235		= NewVectorValue("b_k235", "", "m projection onto psi(Re,Im)_k235", GetModeAmplitudeReImk235 )
	b_k236		= NewVectorValue("b_k236", "", "m projection onto psi(Re,Im)_k236", GetModeAmplitudeReImk236 )
	b_k237		= NewVectorValue("b_k237", "", "m projection onto psi(Re,Im)_k237", GetModeAmplitudeReImk237 )
	b_k238		= NewVectorValue("b_k238", "", "m projection onto psi(Re,Im)_k238", GetModeAmplitudeReImk238 )
	b_k239		= NewVectorValue("b_k239", "", "m projection onto psi(Re,Im)_k239", GetModeAmplitudeReImk239 )
	b_k240		= NewVectorValue("b_k240", "", "m projection onto psi(Re,Im)_k240", GetModeAmplitudeReImk240 )
	b_k241		= NewVectorValue("b_k241", "", "m projection onto psi(Re,Im)_k241", GetModeAmplitudeReImk241 )
	b_k242		= NewVectorValue("b_k242", "", "m projection onto psi(Re,Im)_k242", GetModeAmplitudeReImk242 )
	b_k243		= NewVectorValue("b_k243", "", "m projection onto psi(Re,Im)_k243", GetModeAmplitudeReImk243 )
	b_k244		= NewVectorValue("b_k244", "", "m projection onto psi(Re,Im)_k244", GetModeAmplitudeReImk244 )
	b_k245		= NewVectorValue("b_k245", "", "m projection onto psi(Re,Im)_k245", GetModeAmplitudeReImk245 )
	b_k246		= NewVectorValue("b_k246", "", "m projection onto psi(Re,Im)_k246", GetModeAmplitudeReImk246 )
	b_k247		= NewVectorValue("b_k247", "", "m projection onto psi(Re,Im)_k247", GetModeAmplitudeReImk247 )
	b_k248		= NewVectorValue("b_k248", "", "m projection onto psi(Re,Im)_k248", GetModeAmplitudeReImk248 )
	b_k249		= NewVectorValue("b_k249", "", "m projection onto psi(Re,Im)_k249", GetModeAmplitudeReImk249 )
	b_k250		= NewVectorValue("b_k250", "", "m projection onto psi(Re,Im)_k250", GetModeAmplitudeReImk250 )
	b_k251		= NewVectorValue("b_k251", "", "m projection onto psi(Re,Im)_k251", GetModeAmplitudeReImk251 )
	b_k252		= NewVectorValue("b_k252", "", "m projection onto psi(Re,Im)_k252", GetModeAmplitudeReImk252 )
	b_k253		= NewVectorValue("b_k253", "", "m projection onto psi(Re,Im)_k253", GetModeAmplitudeReImk253 )
	b_k254		= NewVectorValue("b_k254", "", "m projection onto psi(Re,Im)_k254", GetModeAmplitudeReImk254 )
	b_k255		= NewVectorValue("b_k255", "", "m projection onto psi(Re,Im)_k255", GetModeAmplitudeReImk255 )
	b_k256		= NewVectorValue("b_k256", "", "m projection onto psi(Re,Im)_k256", GetModeAmplitudeReImk256 )
	b_k257		= NewVectorValue("b_k257", "", "m projection onto psi(Re,Im)_k257", GetModeAmplitudeReImk257 )
	b_k258		= NewVectorValue("b_k258", "", "m projection onto psi(Re,Im)_k258", GetModeAmplitudeReImk258 )
	b_k259		= NewVectorValue("b_k259", "", "m projection onto psi(Re,Im)_k259", GetModeAmplitudeReImk259 )
	b_k260		= NewVectorValue("b_k260", "", "m projection onto psi(Re,Im)_k260", GetModeAmplitudeReImk260 )
	b_k261		= NewVectorValue("b_k261", "", "m projection onto psi(Re,Im)_k261", GetModeAmplitudeReImk261 )
	b_k262		= NewVectorValue("b_k262", "", "m projection onto psi(Re,Im)_k262", GetModeAmplitudeReImk262 )
	b_k263		= NewVectorValue("b_k263", "", "m projection onto psi(Re,Im)_k263", GetModeAmplitudeReImk263 )
	b_k264		= NewVectorValue("b_k264", "", "m projection onto psi(Re,Im)_k264", GetModeAmplitudeReImk264 )
	b_k265		= NewVectorValue("b_k265", "", "m projection onto psi(Re,Im)_k265", GetModeAmplitudeReImk265 )
	b_k266		= NewVectorValue("b_k266", "", "m projection onto psi(Re,Im)_k266", GetModeAmplitudeReImk266 )
	b_k267		= NewVectorValue("b_k267", "", "m projection onto psi(Re,Im)_k267", GetModeAmplitudeReImk267 )
	b_k268		= NewVectorValue("b_k268", "", "m projection onto psi(Re,Im)_k268", GetModeAmplitudeReImk268 )
	b_k269		= NewVectorValue("b_k269", "", "m projection onto psi(Re,Im)_k269", GetModeAmplitudeReImk269 )
	b_k270		= NewVectorValue("b_k270", "", "m projection onto psi(Re,Im)_k270", GetModeAmplitudeReImk270 )
	b_k271		= NewVectorValue("b_k271", "", "m projection onto psi(Re,Im)_k271", GetModeAmplitudeReImk271 )
	b_k272		= NewVectorValue("b_k272", "", "m projection onto psi(Re,Im)_k272", GetModeAmplitudeReImk272 )
	b_k273		= NewVectorValue("b_k273", "", "m projection onto psi(Re,Im)_k273", GetModeAmplitudeReImk273 )
	b_k274		= NewVectorValue("b_k274", "", "m projection onto psi(Re,Im)_k274", GetModeAmplitudeReImk274 )
	b_k275		= NewVectorValue("b_k275", "", "m projection onto psi(Re,Im)_k275", GetModeAmplitudeReImk275 )
	b_k276		= NewVectorValue("b_k276", "", "m projection onto psi(Re,Im)_k276", GetModeAmplitudeReImk276 )
	b_k277		= NewVectorValue("b_k277", "", "m projection onto psi(Re,Im)_k277", GetModeAmplitudeReImk277 )
	b_k278		= NewVectorValue("b_k278", "", "m projection onto psi(Re,Im)_k278", GetModeAmplitudeReImk278 )
	b_k279		= NewVectorValue("b_k279", "", "m projection onto psi(Re,Im)_k279", GetModeAmplitudeReImk279 )
	b_k280		= NewVectorValue("b_k280", "", "m projection onto psi(Re,Im)_k280", GetModeAmplitudeReImk280 )
	b_k281		= NewVectorValue("b_k281", "", "m projection onto psi(Re,Im)_k281", GetModeAmplitudeReImk281 )
	b_k282		= NewVectorValue("b_k282", "", "m projection onto psi(Re,Im)_k282", GetModeAmplitudeReImk282 )
	b_k283		= NewVectorValue("b_k283", "", "m projection onto psi(Re,Im)_k283", GetModeAmplitudeReImk283 )
	b_k284		= NewVectorValue("b_k284", "", "m projection onto psi(Re,Im)_k284", GetModeAmplitudeReImk284 )
	b_k285		= NewVectorValue("b_k285", "", "m projection onto psi(Re,Im)_k285", GetModeAmplitudeReImk285 )
	b_k286		= NewVectorValue("b_k286", "", "m projection onto psi(Re,Im)_k286", GetModeAmplitudeReImk286 )
	b_k287		= NewVectorValue("b_k287", "", "m projection onto psi(Re,Im)_k287", GetModeAmplitudeReImk287 )
	b_k288		= NewVectorValue("b_k288", "", "m projection onto psi(Re,Im)_k288", GetModeAmplitudeReImk288 )
	b_k289		= NewVectorValue("b_k289", "", "m projection onto psi(Re,Im)_k289", GetModeAmplitudeReImk289 )
	b_k290		= NewVectorValue("b_k290", "", "m projection onto psi(Re,Im)_k290", GetModeAmplitudeReImk290 )
	b_k291		= NewVectorValue("b_k291", "", "m projection onto psi(Re,Im)_k291", GetModeAmplitudeReImk291 )
	b_k292		= NewVectorValue("b_k292", "", "m projection onto psi(Re,Im)_k292", GetModeAmplitudeReImk292 )
	b_k293		= NewVectorValue("b_k293", "", "m projection onto psi(Re,Im)_k293", GetModeAmplitudeReImk293 )
	b_k294		= NewVectorValue("b_k294", "", "m projection onto psi(Re,Im)_k294", GetModeAmplitudeReImk294 )
	b_k295		= NewVectorValue("b_k295", "", "m projection onto psi(Re,Im)_k295", GetModeAmplitudeReImk295 )
	b_k296		= NewVectorValue("b_k296", "", "m projection onto psi(Re,Im)_k296", GetModeAmplitudeReImk296 )
	b_k297		= NewVectorValue("b_k297", "", "m projection onto psi(Re,Im)_k297", GetModeAmplitudeReImk297 )
	b_k298		= NewVectorValue("b_k298", "", "m projection onto psi(Re,Im)_k298", GetModeAmplitudeReImk298 )
	b_k299		= NewVectorValue("b_k299", "", "m projection onto psi(Re,Im)_k299", GetModeAmplitudeReImk299 )
	b_k300		= NewVectorValue("b_k300", "", "m projection onto psi(Re,Im)_k300", GetModeAmplitudeReImk300 )
	b_k301		= NewVectorValue("b_k301", "", "m projection onto psi(Re,Im)_k301", GetModeAmplitudeReImk301 )
	b_k302		= NewVectorValue("b_k302", "", "m projection onto psi(Re,Im)_k302", GetModeAmplitudeReImk302 )
	b_k303		= NewVectorValue("b_k303", "", "m projection onto psi(Re,Im)_k303", GetModeAmplitudeReImk303 )
	b_k304		= NewVectorValue("b_k304", "", "m projection onto psi(Re,Im)_k304", GetModeAmplitudeReImk304 )
	b_k305		= NewVectorValue("b_k305", "", "m projection onto psi(Re,Im)_k305", GetModeAmplitudeReImk305 )
	b_k306		= NewVectorValue("b_k306", "", "m projection onto psi(Re,Im)_k306", GetModeAmplitudeReImk306 )
	b_k307		= NewVectorValue("b_k307", "", "m projection onto psi(Re,Im)_k307", GetModeAmplitudeReImk307 )
	b_k308		= NewVectorValue("b_k308", "", "m projection onto psi(Re,Im)_k308", GetModeAmplitudeReImk308 )
	b_k309		= NewVectorValue("b_k309", "", "m projection onto psi(Re,Im)_k309", GetModeAmplitudeReImk309 )
	b_k310		= NewVectorValue("b_k310", "", "m projection onto psi(Re,Im)_k310", GetModeAmplitudeReImk310 )
	b_k311		= NewVectorValue("b_k311", "", "m projection onto psi(Re,Im)_k311", GetModeAmplitudeReImk311 )
	b_k312		= NewVectorValue("b_k312", "", "m projection onto psi(Re,Im)_k312", GetModeAmplitudeReImk312 )
	b_k313		= NewVectorValue("b_k313", "", "m projection onto psi(Re,Im)_k313", GetModeAmplitudeReImk313 )
	b_k314		= NewVectorValue("b_k314", "", "m projection onto psi(Re,Im)_k314", GetModeAmplitudeReImk314 )
	b_k315		= NewVectorValue("b_k315", "", "m projection onto psi(Re,Im)_k315", GetModeAmplitudeReImk315 )
	b_k316		= NewVectorValue("b_k316", "", "m projection onto psi(Re,Im)_k316", GetModeAmplitudeReImk316 )
	b_k317		= NewVectorValue("b_k317", "", "m projection onto psi(Re,Im)_k317", GetModeAmplitudeReImk317 )
	b_k318		= NewVectorValue("b_k318", "", "m projection onto psi(Re,Im)_k318", GetModeAmplitudeReImk318 )
	b_k319		= NewVectorValue("b_k319", "", "m projection onto psi(Re,Im)_k319", GetModeAmplitudeReImk319 )
	b_k320		= NewVectorValue("b_k320", "", "m projection onto psi(Re,Im)_k320", GetModeAmplitudeReImk320 )
	b_k321		= NewVectorValue("b_k321", "", "m projection onto psi(Re,Im)_k321", GetModeAmplitudeReImk321 )
	b_k322		= NewVectorValue("b_k322", "", "m projection onto psi(Re,Im)_k322", GetModeAmplitudeReImk322 )
	b_k323		= NewVectorValue("b_k323", "", "m projection onto psi(Re,Im)_k323", GetModeAmplitudeReImk323 )
	b_k324		= NewVectorValue("b_k324", "", "m projection onto psi(Re,Im)_k324", GetModeAmplitudeReImk324 )
	b_k325		= NewVectorValue("b_k325", "", "m projection onto psi(Re,Im)_k325", GetModeAmplitudeReImk325 )
	b_k326		= NewVectorValue("b_k326", "", "m projection onto psi(Re,Im)_k326", GetModeAmplitudeReImk326 )
	b_k327		= NewVectorValue("b_k327", "", "m projection onto psi(Re,Im)_k327", GetModeAmplitudeReImk327 )
	b_k328		= NewVectorValue("b_k328", "", "m projection onto psi(Re,Im)_k328", GetModeAmplitudeReImk328 )
	b_k329		= NewVectorValue("b_k329", "", "m projection onto psi(Re,Im)_k329", GetModeAmplitudeReImk329 )
	b_k330		= NewVectorValue("b_k330", "", "m projection onto psi(Re,Im)_k330", GetModeAmplitudeReImk330 )
	b_k331		= NewVectorValue("b_k331", "", "m projection onto psi(Re,Im)_k331", GetModeAmplitudeReImk331 )
	b_k332		= NewVectorValue("b_k332", "", "m projection onto psi(Re,Im)_k332", GetModeAmplitudeReImk332 )
	b_k333		= NewVectorValue("b_k333", "", "m projection onto psi(Re,Im)_k333", GetModeAmplitudeReImk333 )
	b_k334		= NewVectorValue("b_k334", "", "m projection onto psi(Re,Im)_k334", GetModeAmplitudeReImk334 )
	b_k335		= NewVectorValue("b_k335", "", "m projection onto psi(Re,Im)_k335", GetModeAmplitudeReImk335 )
	b_k336		= NewVectorValue("b_k336", "", "m projection onto psi(Re,Im)_k336", GetModeAmplitudeReImk336 )
	b_k337		= NewVectorValue("b_k337", "", "m projection onto psi(Re,Im)_k337", GetModeAmplitudeReImk337 )
	b_k338		= NewVectorValue("b_k338", "", "m projection onto psi(Re,Im)_k338", GetModeAmplitudeReImk338 )
	b_k339		= NewVectorValue("b_k339", "", "m projection onto psi(Re,Im)_k339", GetModeAmplitudeReImk339 )
	b_k340		= NewVectorValue("b_k340", "", "m projection onto psi(Re,Im)_k340", GetModeAmplitudeReImk340 )
	b_k341		= NewVectorValue("b_k341", "", "m projection onto psi(Re,Im)_k341", GetModeAmplitudeReImk341 )
	b_k342		= NewVectorValue("b_k342", "", "m projection onto psi(Re,Im)_k342", GetModeAmplitudeReImk342 )
	b_k343		= NewVectorValue("b_k343", "", "m projection onto psi(Re,Im)_k343", GetModeAmplitudeReImk343 )
	b_k344		= NewVectorValue("b_k344", "", "m projection onto psi(Re,Im)_k344", GetModeAmplitudeReImk344 )
	b_k345		= NewVectorValue("b_k345", "", "m projection onto psi(Re,Im)_k345", GetModeAmplitudeReImk345 )
	b_k346		= NewVectorValue("b_k346", "", "m projection onto psi(Re,Im)_k346", GetModeAmplitudeReImk346 )
	b_k347		= NewVectorValue("b_k347", "", "m projection onto psi(Re,Im)_k347", GetModeAmplitudeReImk347 )
	b_k348		= NewVectorValue("b_k348", "", "m projection onto psi(Re,Im)_k348", GetModeAmplitudeReImk348 )
	b_k349		= NewVectorValue("b_k349", "", "m projection onto psi(Re,Im)_k349", GetModeAmplitudeReImk349 )
	b_k350		= NewVectorValue("b_k350", "", "m projection onto psi(Re,Im)_k350", GetModeAmplitudeReImk350 )
	b_k351		= NewVectorValue("b_k351", "", "m projection onto psi(Re,Im)_k351", GetModeAmplitudeReImk351 )
	b_k352		= NewVectorValue("b_k352", "", "m projection onto psi(Re,Im)_k352", GetModeAmplitudeReImk352 )
	b_k353		= NewVectorValue("b_k353", "", "m projection onto psi(Re,Im)_k353", GetModeAmplitudeReImk353 )
	b_k354		= NewVectorValue("b_k354", "", "m projection onto psi(Re,Im)_k354", GetModeAmplitudeReImk354 )
	b_k355		= NewVectorValue("b_k355", "", "m projection onto psi(Re,Im)_k355", GetModeAmplitudeReImk355 )
	b_k356		= NewVectorValue("b_k356", "", "m projection onto psi(Re,Im)_k356", GetModeAmplitudeReImk356 )
	b_k357		= NewVectorValue("b_k357", "", "m projection onto psi(Re,Im)_k357", GetModeAmplitudeReImk357 )
	b_k358		= NewVectorValue("b_k358", "", "m projection onto psi(Re,Im)_k358", GetModeAmplitudeReImk358 )
	b_k359		= NewVectorValue("b_k359", "", "m projection onto psi(Re,Im)_k359", GetModeAmplitudeReImk359 )
	b_k360		= NewVectorValue("b_k360", "", "m projection onto psi(Re,Im)_k360", GetModeAmplitudeReImk360 )
	b_k361		= NewVectorValue("b_k361", "", "m projection onto psi(Re,Im)_k361", GetModeAmplitudeReImk361 )
	b_k362		= NewVectorValue("b_k362", "", "m projection onto psi(Re,Im)_k362", GetModeAmplitudeReImk362 )
	b_k363		= NewVectorValue("b_k363", "", "m projection onto psi(Re,Im)_k363", GetModeAmplitudeReImk363 )
	b_k364		= NewVectorValue("b_k364", "", "m projection onto psi(Re,Im)_k364", GetModeAmplitudeReImk364 )
	b_k365		= NewVectorValue("b_k365", "", "m projection onto psi(Re,Im)_k365", GetModeAmplitudeReImk365 )
	b_k366		= NewVectorValue("b_k366", "", "m projection onto psi(Re,Im)_k366", GetModeAmplitudeReImk366 )
	b_k367		= NewVectorValue("b_k367", "", "m projection onto psi(Re,Im)_k367", GetModeAmplitudeReImk367 )
	b_k368		= NewVectorValue("b_k368", "", "m projection onto psi(Re,Im)_k368", GetModeAmplitudeReImk368 )
	b_k369		= NewVectorValue("b_k369", "", "m projection onto psi(Re,Im)_k369", GetModeAmplitudeReImk369 )
	b_k370		= NewVectorValue("b_k370", "", "m projection onto psi(Re,Im)_k370", GetModeAmplitudeReImk370 )
	b_k371		= NewVectorValue("b_k371", "", "m projection onto psi(Re,Im)_k371", GetModeAmplitudeReImk371 )
	b_k372		= NewVectorValue("b_k372", "", "m projection onto psi(Re,Im)_k372", GetModeAmplitudeReImk372 )
	b_k373		= NewVectorValue("b_k373", "", "m projection onto psi(Re,Im)_k373", GetModeAmplitudeReImk373 )
	b_k374		= NewVectorValue("b_k374", "", "m projection onto psi(Re,Im)_k374", GetModeAmplitudeReImk374 )
	b_k375		= NewVectorValue("b_k375", "", "m projection onto psi(Re,Im)_k375", GetModeAmplitudeReImk375 )
	b_k376		= NewVectorValue("b_k376", "", "m projection onto psi(Re,Im)_k376", GetModeAmplitudeReImk376 )
	b_k377		= NewVectorValue("b_k377", "", "m projection onto psi(Re,Im)_k377", GetModeAmplitudeReImk377 )
	b_k378		= NewVectorValue("b_k378", "", "m projection onto psi(Re,Im)_k378", GetModeAmplitudeReImk378 )
	b_k379		= NewVectorValue("b_k379", "", "m projection onto psi(Re,Im)_k379", GetModeAmplitudeReImk379 )
	b_k380		= NewVectorValue("b_k380", "", "m projection onto psi(Re,Im)_k380", GetModeAmplitudeReImk380 )
	b_k381		= NewVectorValue("b_k381", "", "m projection onto psi(Re,Im)_k381", GetModeAmplitudeReImk381 )
	b_k382		= NewVectorValue("b_k382", "", "m projection onto psi(Re,Im)_k382", GetModeAmplitudeReImk382 )
	b_k383		= NewVectorValue("b_k383", "", "m projection onto psi(Re,Im)_k383", GetModeAmplitudeReImk383 )
	b_k384		= NewVectorValue("b_k384", "", "m projection onto psi(Re,Im)_k384", GetModeAmplitudeReImk384 )
	b_k385		= NewVectorValue("b_k385", "", "m projection onto psi(Re,Im)_k385", GetModeAmplitudeReImk385 )
	b_k386		= NewVectorValue("b_k386", "", "m projection onto psi(Re,Im)_k386", GetModeAmplitudeReImk386 )
	b_k387		= NewVectorValue("b_k387", "", "m projection onto psi(Re,Im)_k387", GetModeAmplitudeReImk387 )
	b_k388		= NewVectorValue("b_k388", "", "m projection onto psi(Re,Im)_k388", GetModeAmplitudeReImk388 )
	b_k389		= NewVectorValue("b_k389", "", "m projection onto psi(Re,Im)_k389", GetModeAmplitudeReImk389 )
	b_k390		= NewVectorValue("b_k390", "", "m projection onto psi(Re,Im)_k390", GetModeAmplitudeReImk390 )
	b_k391		= NewVectorValue("b_k391", "", "m projection onto psi(Re,Im)_k391", GetModeAmplitudeReImk391 )
	b_k392		= NewVectorValue("b_k392", "", "m projection onto psi(Re,Im)_k392", GetModeAmplitudeReImk392 )
	b_k393		= NewVectorValue("b_k393", "", "m projection onto psi(Re,Im)_k393", GetModeAmplitudeReImk393 )
	b_k394		= NewVectorValue("b_k394", "", "m projection onto psi(Re,Im)_k394", GetModeAmplitudeReImk394 )
	b_k395		= NewVectorValue("b_k395", "", "m projection onto psi(Re,Im)_k395", GetModeAmplitudeReImk395 )
	b_k396		= NewVectorValue("b_k396", "", "m projection onto psi(Re,Im)_k396", GetModeAmplitudeReImk396 )
	b_k397		= NewVectorValue("b_k397", "", "m projection onto psi(Re,Im)_k397", GetModeAmplitudeReImk397 )
	b_k398		= NewVectorValue("b_k398", "", "m projection onto psi(Re,Im)_k398", GetModeAmplitudeReImk398 )
	b_k399		= NewVectorValue("b_k399", "", "m projection onto psi(Re,Im)_k399", GetModeAmplitudeReImk399 )
	b_k400		= NewVectorValue("b_k400", "", "m projection onto psi(Re,Im)_k400", GetModeAmplitudeReImk400 )
	b_k401		= NewVectorValue("b_k401", "", "m projection onto psi(Re,Im)_k401", GetModeAmplitudeReImk401 )
	b_k402		= NewVectorValue("b_k402", "", "m projection onto psi(Re,Im)_k402", GetModeAmplitudeReImk402 )
	b_k403		= NewVectorValue("b_k403", "", "m projection onto psi(Re,Im)_k403", GetModeAmplitudeReImk403 )
	b_k404		= NewVectorValue("b_k404", "", "m projection onto psi(Re,Im)_k404", GetModeAmplitudeReImk404 )
	b_k405		= NewVectorValue("b_k405", "", "m projection onto psi(Re,Im)_k405", GetModeAmplitudeReImk405 )
	b_k406		= NewVectorValue("b_k406", "", "m projection onto psi(Re,Im)_k406", GetModeAmplitudeReImk406 )
	b_k407		= NewVectorValue("b_k407", "", "m projection onto psi(Re,Im)_k407", GetModeAmplitudeReImk407 )
	b_k408		= NewVectorValue("b_k408", "", "m projection onto psi(Re,Im)_k408", GetModeAmplitudeReImk408 )
	b_k409		= NewVectorValue("b_k409", "", "m projection onto psi(Re,Im)_k409", GetModeAmplitudeReImk409 )
	b_k410		= NewVectorValue("b_k410", "", "m projection onto psi(Re,Im)_k410", GetModeAmplitudeReImk410 )
	b_k411		= NewVectorValue("b_k411", "", "m projection onto psi(Re,Im)_k411", GetModeAmplitudeReImk411 )
	b_k412		= NewVectorValue("b_k412", "", "m projection onto psi(Re,Im)_k412", GetModeAmplitudeReImk412 )
	b_k413		= NewVectorValue("b_k413", "", "m projection onto psi(Re,Im)_k413", GetModeAmplitudeReImk413 )
	b_k414		= NewVectorValue("b_k414", "", "m projection onto psi(Re,Im)_k414", GetModeAmplitudeReImk414 )
	b_k415		= NewVectorValue("b_k415", "", "m projection onto psi(Re,Im)_k415", GetModeAmplitudeReImk415 )
	b_k416		= NewVectorValue("b_k416", "", "m projection onto psi(Re,Im)_k416", GetModeAmplitudeReImk416 )
	b_k417		= NewVectorValue("b_k417", "", "m projection onto psi(Re,Im)_k417", GetModeAmplitudeReImk417 )
	b_k418		= NewVectorValue("b_k418", "", "m projection onto psi(Re,Im)_k418", GetModeAmplitudeReImk418 )
	b_k419		= NewVectorValue("b_k419", "", "m projection onto psi(Re,Im)_k419", GetModeAmplitudeReImk419 )
	b_k420		= NewVectorValue("b_k420", "", "m projection onto psi(Re,Im)_k420", GetModeAmplitudeReImk420 )
	b_k421		= NewVectorValue("b_k421", "", "m projection onto psi(Re,Im)_k421", GetModeAmplitudeReImk421 )
	b_k422		= NewVectorValue("b_k422", "", "m projection onto psi(Re,Im)_k422", GetModeAmplitudeReImk422 )
	b_k423		= NewVectorValue("b_k423", "", "m projection onto psi(Re,Im)_k423", GetModeAmplitudeReImk423 )
	b_k424		= NewVectorValue("b_k424", "", "m projection onto psi(Re,Im)_k424", GetModeAmplitudeReImk424 )
	b_k425		= NewVectorValue("b_k425", "", "m projection onto psi(Re,Im)_k425", GetModeAmplitudeReImk425 )
	b_k426		= NewVectorValue("b_k426", "", "m projection onto psi(Re,Im)_k426", GetModeAmplitudeReImk426 )
	b_k427		= NewVectorValue("b_k427", "", "m projection onto psi(Re,Im)_k427", GetModeAmplitudeReImk427 )
	b_k428		= NewVectorValue("b_k428", "", "m projection onto psi(Re,Im)_k428", GetModeAmplitudeReImk428 )
	b_k429		= NewVectorValue("b_k429", "", "m projection onto psi(Re,Im)_k429", GetModeAmplitudeReImk429 )
	b_k430		= NewVectorValue("b_k430", "", "m projection onto psi(Re,Im)_k430", GetModeAmplitudeReImk430 )
	b_k431		= NewVectorValue("b_k431", "", "m projection onto psi(Re,Im)_k431", GetModeAmplitudeReImk431 )
	b_k432		= NewVectorValue("b_k432", "", "m projection onto psi(Re,Im)_k432", GetModeAmplitudeReImk432 )
	b_k433		= NewVectorValue("b_k433", "", "m projection onto psi(Re,Im)_k433", GetModeAmplitudeReImk433 )
	b_k434		= NewVectorValue("b_k434", "", "m projection onto psi(Re,Im)_k434", GetModeAmplitudeReImk434 )
	b_k435		= NewVectorValue("b_k435", "", "m projection onto psi(Re,Im)_k435", GetModeAmplitudeReImk435 )
	b_k436		= NewVectorValue("b_k436", "", "m projection onto psi(Re,Im)_k436", GetModeAmplitudeReImk436 )
	b_k437		= NewVectorValue("b_k437", "", "m projection onto psi(Re,Im)_k437", GetModeAmplitudeReImk437 )
	b_k438		= NewVectorValue("b_k438", "", "m projection onto psi(Re,Im)_k438", GetModeAmplitudeReImk438 )
	b_k439		= NewVectorValue("b_k439", "", "m projection onto psi(Re,Im)_k439", GetModeAmplitudeReImk439 )
	b_k440		= NewVectorValue("b_k440", "", "m projection onto psi(Re,Im)_k440", GetModeAmplitudeReImk440 )
	b_k441		= NewVectorValue("b_k441", "", "m projection onto psi(Re,Im)_k441", GetModeAmplitudeReImk441 )
	b_k442		= NewVectorValue("b_k442", "", "m projection onto psi(Re,Im)_k442", GetModeAmplitudeReImk442 )
	b_k443		= NewVectorValue("b_k443", "", "m projection onto psi(Re,Im)_k443", GetModeAmplitudeReImk443 )
	b_k444		= NewVectorValue("b_k444", "", "m projection onto psi(Re,Im)_k444", GetModeAmplitudeReImk444 )
	b_k445		= NewVectorValue("b_k445", "", "m projection onto psi(Re,Im)_k445", GetModeAmplitudeReImk445 )
	b_k446		= NewVectorValue("b_k446", "", "m projection onto psi(Re,Im)_k446", GetModeAmplitudeReImk446 )
	b_k447		= NewVectorValue("b_k447", "", "m projection onto psi(Re,Im)_k447", GetModeAmplitudeReImk447 )
	b_k448		= NewVectorValue("b_k448", "", "m projection onto psi(Re,Im)_k448", GetModeAmplitudeReImk448 )
	b_k449		= NewVectorValue("b_k449", "", "m projection onto psi(Re,Im)_k449", GetModeAmplitudeReImk449 )
	b_k450		= NewVectorValue("b_k450", "", "m projection onto psi(Re,Im)_k450", GetModeAmplitudeReImk450 )
	b_k451		= NewVectorValue("b_k451", "", "m projection onto psi(Re,Im)_k451", GetModeAmplitudeReImk451 )
	b_k452		= NewVectorValue("b_k452", "", "m projection onto psi(Re,Im)_k452", GetModeAmplitudeReImk452 )
	b_k453		= NewVectorValue("b_k453", "", "m projection onto psi(Re,Im)_k453", GetModeAmplitudeReImk453 )
	b_k454		= NewVectorValue("b_k454", "", "m projection onto psi(Re,Im)_k454", GetModeAmplitudeReImk454 )
	b_k455		= NewVectorValue("b_k455", "", "m projection onto psi(Re,Im)_k455", GetModeAmplitudeReImk455 )
	b_k456		= NewVectorValue("b_k456", "", "m projection onto psi(Re,Im)_k456", GetModeAmplitudeReImk456 )
	b_k457		= NewVectorValue("b_k457", "", "m projection onto psi(Re,Im)_k457", GetModeAmplitudeReImk457 )
	b_k458		= NewVectorValue("b_k458", "", "m projection onto psi(Re,Im)_k458", GetModeAmplitudeReImk458 )
	b_k459		= NewVectorValue("b_k459", "", "m projection onto psi(Re,Im)_k459", GetModeAmplitudeReImk459 )
	b_k460		= NewVectorValue("b_k460", "", "m projection onto psi(Re,Im)_k460", GetModeAmplitudeReImk460 )
	b_k461		= NewVectorValue("b_k461", "", "m projection onto psi(Re,Im)_k461", GetModeAmplitudeReImk461 )
	b_k462		= NewVectorValue("b_k462", "", "m projection onto psi(Re,Im)_k462", GetModeAmplitudeReImk462 )
	b_k463		= NewVectorValue("b_k463", "", "m projection onto psi(Re,Im)_k463", GetModeAmplitudeReImk463 )
	b_k464		= NewVectorValue("b_k464", "", "m projection onto psi(Re,Im)_k464", GetModeAmplitudeReImk464 )
	b_k465		= NewVectorValue("b_k465", "", "m projection onto psi(Re,Im)_k465", GetModeAmplitudeReImk465 )
	b_k466		= NewVectorValue("b_k466", "", "m projection onto psi(Re,Im)_k466", GetModeAmplitudeReImk466 )
	b_k467		= NewVectorValue("b_k467", "", "m projection onto psi(Re,Im)_k467", GetModeAmplitudeReImk467 )
	b_k468		= NewVectorValue("b_k468", "", "m projection onto psi(Re,Im)_k468", GetModeAmplitudeReImk468 )
	b_k469		= NewVectorValue("b_k469", "", "m projection onto psi(Re,Im)_k469", GetModeAmplitudeReImk469 )
	b_k470		= NewVectorValue("b_k470", "", "m projection onto psi(Re,Im)_k470", GetModeAmplitudeReImk470 )
	b_k471		= NewVectorValue("b_k471", "", "m projection onto psi(Re,Im)_k471", GetModeAmplitudeReImk471 )
	b_k472		= NewVectorValue("b_k472", "", "m projection onto psi(Re,Im)_k472", GetModeAmplitudeReImk472 )
	b_k473		= NewVectorValue("b_k473", "", "m projection onto psi(Re,Im)_k473", GetModeAmplitudeReImk473 )
	b_k474		= NewVectorValue("b_k474", "", "m projection onto psi(Re,Im)_k474", GetModeAmplitudeReImk474 )
	b_k475		= NewVectorValue("b_k475", "", "m projection onto psi(Re,Im)_k475", GetModeAmplitudeReImk475 )
	b_k476		= NewVectorValue("b_k476", "", "m projection onto psi(Re,Im)_k476", GetModeAmplitudeReImk476 )
	b_k477		= NewVectorValue("b_k477", "", "m projection onto psi(Re,Im)_k477", GetModeAmplitudeReImk477 )
	b_k478		= NewVectorValue("b_k478", "", "m projection onto psi(Re,Im)_k478", GetModeAmplitudeReImk478 )
	b_k479		= NewVectorValue("b_k479", "", "m projection onto psi(Re,Im)_k479", GetModeAmplitudeReImk479 )
	b_k480		= NewVectorValue("b_k480", "", "m projection onto psi(Re,Im)_k480", GetModeAmplitudeReImk480 )
	b_k481		= NewVectorValue("b_k481", "", "m projection onto psi(Re,Im)_k481", GetModeAmplitudeReImk481 )
	b_k482		= NewVectorValue("b_k482", "", "m projection onto psi(Re,Im)_k482", GetModeAmplitudeReImk482 )
	b_k483		= NewVectorValue("b_k483", "", "m projection onto psi(Re,Im)_k483", GetModeAmplitudeReImk483 )
	b_k484		= NewVectorValue("b_k484", "", "m projection onto psi(Re,Im)_k484", GetModeAmplitudeReImk484 )
	b_k485		= NewVectorValue("b_k485", "", "m projection onto psi(Re,Im)_k485", GetModeAmplitudeReImk485 )
	b_k486		= NewVectorValue("b_k486", "", "m projection onto psi(Re,Im)_k486", GetModeAmplitudeReImk486 )
	b_k487		= NewVectorValue("b_k487", "", "m projection onto psi(Re,Im)_k487", GetModeAmplitudeReImk487 )
	b_k488		= NewVectorValue("b_k488", "", "m projection onto psi(Re,Im)_k488", GetModeAmplitudeReImk488 )
	b_k489		= NewVectorValue("b_k489", "", "m projection onto psi(Re,Im)_k489", GetModeAmplitudeReImk489 )
	b_k490		= NewVectorValue("b_k490", "", "m projection onto psi(Re,Im)_k490", GetModeAmplitudeReImk490 )
	b_k491		= NewVectorValue("b_k491", "", "m projection onto psi(Re,Im)_k491", GetModeAmplitudeReImk491 )
	b_k492		= NewVectorValue("b_k492", "", "m projection onto psi(Re,Im)_k492", GetModeAmplitudeReImk492 )
	b_k493		= NewVectorValue("b_k493", "", "m projection onto psi(Re,Im)_k493", GetModeAmplitudeReImk493 )
	b_k494		= NewVectorValue("b_k494", "", "m projection onto psi(Re,Im)_k494", GetModeAmplitudeReImk494 )
	b_k495		= NewVectorValue("b_k495", "", "m projection onto psi(Re,Im)_k495", GetModeAmplitudeReImk495 )
	b_k496		= NewVectorValue("b_k496", "", "m projection onto psi(Re,Im)_k496", GetModeAmplitudeReImk496 )
	b_k497		= NewVectorValue("b_k497", "", "m projection onto psi(Re,Im)_k497", GetModeAmplitudeReImk497 )
	b_k498		= NewVectorValue("b_k498", "", "m projection onto psi(Re,Im)_k498", GetModeAmplitudeReImk498 )
	b_k499		= NewVectorValue("b_k499", "", "m projection onto psi(Re,Im)_k499", GetModeAmplitudeReImk499 )
)
	
	
func GetModeAmplitudeReImk000() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k000), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k000), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk001() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k001), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k001), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk002() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k002), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k002), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk003() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k003), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k003), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk004() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k004), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k004), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk005() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k005), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k005), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk006() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k006), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k006), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk007() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k007), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k007), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk008() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k008), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k008), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk009() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k009), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k009), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk010() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k010), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k010), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk011() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k011), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k011), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk012() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k012), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k012), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk013() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k013), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k013), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk014() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k014), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k014), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk015() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k015), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k015), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk016() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k016), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k016), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk017() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k017), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k017), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk018() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k018), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k018), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk019() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k019), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k019), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk020() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k020), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k020), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk021() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k021), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k021), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk022() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k022), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k022), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk023() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k023), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k023), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk024() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k024), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k024), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk025() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k025), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k025), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk026() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k026), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k026), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk027() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k027), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k027), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk028() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k028), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k028), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk029() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k029), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k029), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk030() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k030), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k030), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk031() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k031), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k031), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk032() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k032), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k032), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk033() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k033), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k033), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk034() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k034), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k034), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk035() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k035), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k035), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk036() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k036), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k036), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk037() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k037), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k037), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk038() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k038), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k038), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk039() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k039), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k039), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk040() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k040), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k040), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk041() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k041), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k041), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk042() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k042), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k042), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk043() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k043), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k043), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk044() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k044), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k044), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk045() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k045), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k045), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk046() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k046), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k046), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk047() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k047), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k047), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk048() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k048), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k048), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk049() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k049), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k049), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk050() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k050), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k050), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk051() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k051), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k051), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk052() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k052), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k052), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk053() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k053), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k053), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk054() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k054), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k054), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk055() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k055), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k055), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk056() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k056), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k056), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk057() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k057), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k057), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk058() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k058), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k058), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk059() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k059), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k059), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk060() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k060), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k060), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk061() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k061), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k061), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk062() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k062), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k062), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk063() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k063), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k063), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk064() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k064), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k064), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk065() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k065), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k065), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk066() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k066), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k066), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk067() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k067), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k067), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk068() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k068), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k068), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk069() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k069), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k069), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk070() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k070), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k070), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk071() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k071), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k071), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk072() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k072), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k072), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk073() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k073), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k073), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk074() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k074), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k074), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk075() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k075), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k075), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk076() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k076), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k076), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk077() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k077), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k077), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk078() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k078), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k078), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk079() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k079), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k079), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk080() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k080), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k080), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk081() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k081), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k081), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk082() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k082), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k082), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk083() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k083), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k083), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk084() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k084), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k084), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk085() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k085), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k085), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk086() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k086), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k086), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk087() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k087), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k087), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk088() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k088), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k088), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk089() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k089), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k089), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk090() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k090), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k090), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk091() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k091), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k091), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk092() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k092), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k092), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk093() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k093), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k093), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk094() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k094), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k094), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk095() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k095), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k095), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk096() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k096), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k096), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk097() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k097), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k097), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk098() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k098), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k098), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk099() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k099), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k099), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk100() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k100), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k100), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk101() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k101), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k101), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk102() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k102), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k102), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk103() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k103), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k103), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk104() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k104), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k104), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk105() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k105), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k105), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk106() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k106), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k106), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk107() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k107), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k107), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk108() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k108), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k108), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk109() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k109), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k109), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk110() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k110), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k110), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk111() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k111), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k111), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk112() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k112), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k112), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk113() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k113), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k113), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk114() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k114), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k114), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk115() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k115), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k115), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk116() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k116), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k116), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk117() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k117), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k117), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk118() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k118), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k118), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk119() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k119), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k119), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk120() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k120), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k120), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk121() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k121), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k121), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk122() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k122), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k122), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk123() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k123), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k123), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk124() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k124), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k124), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk125() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k125), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k125), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk126() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k126), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k126), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk127() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k127), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k127), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk128() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k128), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k128), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk129() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k129), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k129), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk130() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k130), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k130), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk131() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k131), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k131), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk132() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k132), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k132), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk133() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k133), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k133), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk134() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k134), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k134), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk135() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k135), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k135), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk136() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k136), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k136), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk137() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k137), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k137), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk138() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k138), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k138), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk139() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k139), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k139), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk140() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k140), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k140), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk141() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k141), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k141), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk142() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k142), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k142), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk143() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k143), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k143), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk144() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k144), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k144), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk145() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k145), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k145), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk146() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k146), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k146), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk147() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k147), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k147), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk148() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k148), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k148), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk149() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k149), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k149), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk150() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k150), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k150), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk151() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k151), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k151), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk152() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k152), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k152), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk153() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k153), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k153), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk154() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k154), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k154), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk155() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k155), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k155), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk156() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k156), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k156), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk157() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k157), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k157), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk158() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k158), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k158), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk159() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k159), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k159), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk160() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k160), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k160), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk161() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k161), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k161), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk162() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k162), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k162), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk163() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k163), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k163), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk164() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k164), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k164), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk165() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k165), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k165), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk166() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k166), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k166), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk167() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k167), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k167), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk168() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k168), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k168), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk169() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k169), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k169), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk170() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k170), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k170), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk171() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k171), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k171), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk172() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k172), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k172), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk173() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k173), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k173), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk174() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k174), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k174), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk175() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k175), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k175), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk176() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k176), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k176), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk177() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k177), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k177), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk178() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k178), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k178), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk179() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k179), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k179), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk180() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k180), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k180), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk181() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k181), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k181), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk182() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k182), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k182), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk183() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k183), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k183), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk184() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k184), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k184), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk185() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k185), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k185), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk186() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k186), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k186), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk187() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k187), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k187), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk188() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k188), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k188), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk189() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k189), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k189), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk190() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k190), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k190), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk191() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k191), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k191), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk192() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k192), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k192), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk193() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k193), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k193), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk194() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k194), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k194), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk195() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k195), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k195), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk196() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k196), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k196), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk197() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k197), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k197), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk198() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k198), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k198), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk199() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k199), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k199), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk200() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k200), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k200), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk201() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k201), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k201), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk202() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k202), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k202), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk203() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k203), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k203), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk204() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k204), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k204), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk205() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k205), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k205), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk206() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k206), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k206), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk207() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k207), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k207), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk208() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k208), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k208), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk209() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k209), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k209), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk210() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k210), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k210), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk211() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k211), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k211), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk212() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k212), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k212), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk213() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k213), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k213), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk214() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k214), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k214), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk215() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k215), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k215), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk216() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k216), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k216), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk217() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k217), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k217), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk218() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k218), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k218), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk219() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k219), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k219), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk220() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k220), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k220), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk221() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k221), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k221), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk222() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k222), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k222), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk223() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k223), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k223), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk224() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k224), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k224), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk225() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k225), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k225), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk226() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k226), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k226), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk227() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k227), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k227), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk228() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k228), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k228), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk229() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k229), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k229), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk230() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k230), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k230), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk231() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k231), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k231), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk232() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k232), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k232), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk233() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k233), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k233), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk234() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k234), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k234), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk235() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k235), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k235), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk236() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k236), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k236), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk237() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k237), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k237), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk238() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k238), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k238), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk239() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k239), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k239), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk240() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k240), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k240), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk241() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k241), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k241), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk242() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k242), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k242), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk243() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k243), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k243), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk244() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k244), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k244), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk245() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k245), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k245), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk246() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k246), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k246), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk247() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k247), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k247), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk248() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k248), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k248), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk249() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k249), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k249), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk250() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k250), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k250), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk251() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k251), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k251), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk252() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k252), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k252), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk253() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k253), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k253), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk254() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k254), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k254), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk255() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k255), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k255), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk256() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k256), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k256), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk257() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k257), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k257), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk258() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k258), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k258), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk259() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k259), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k259), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk260() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k260), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k260), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk261() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k261), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k261), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk262() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k262), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k262), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk263() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k263), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k263), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk264() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k264), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k264), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk265() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k265), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k265), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk266() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k266), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k266), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk267() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k267), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k267), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk268() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k268), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k268), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk269() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k269), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k269), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk270() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k270), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k270), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk271() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k271), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k271), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk272() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k272), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k272), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk273() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k273), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k273), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk274() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k274), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k274), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk275() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k275), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k275), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk276() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k276), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k276), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk277() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k277), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k277), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk278() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k278), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k278), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk279() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k279), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k279), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk280() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k280), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k280), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk281() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k281), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k281), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk282() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k282), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k282), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk283() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k283), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k283), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk284() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k284), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k284), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk285() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k285), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k285), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk286() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k286), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k286), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk287() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k287), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k287), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk288() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k288), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k288), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk289() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k289), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k289), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk290() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k290), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k290), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk291() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k291), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k291), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk292() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k292), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k292), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk293() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k293), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k293), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk294() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k294), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k294), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk295() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k295), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k295), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk296() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k296), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k296), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk297() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k297), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k297), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk298() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k298), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k298), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk299() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k299), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k299), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk300() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k300), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k300), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk301() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k301), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k301), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk302() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k302), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k302), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk303() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k303), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k303), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk304() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k304), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k304), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk305() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k305), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k305), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk306() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k306), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k306), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk307() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k307), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k307), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk308() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k308), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k308), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk309() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k309), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k309), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk310() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k310), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k310), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk311() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k311), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k311), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk312() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k312), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k312), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk313() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k313), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k313), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk314() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k314), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k314), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk315() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k315), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k315), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk316() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k316), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k316), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk317() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k317), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k317), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk318() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k318), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k318), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk319() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k319), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k319), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk320() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k320), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k320), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk321() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k321), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k321), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk322() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k322), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k322), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk323() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k323), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k323), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk324() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k324), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k324), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk325() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k325), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k325), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk326() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k326), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k326), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk327() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k327), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k327), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk328() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k328), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k328), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk329() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k329), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k329), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk330() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k330), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k330), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk331() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k331), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k331), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk332() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k332), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k332), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk333() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k333), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k333), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk334() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k334), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k334), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk335() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k335), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k335), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk336() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k336), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k336), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk337() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k337), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k337), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk338() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k338), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k338), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk339() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k339), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k339), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk340() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k340), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k340), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk341() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k341), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k341), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk342() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k342), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k342), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk343() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k343), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k343), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk344() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k344), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k344), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk345() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k345), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k345), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk346() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k346), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k346), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk347() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k347), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k347), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk348() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k348), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k348), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk349() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k349), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k349), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk350() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k350), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k350), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk351() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k351), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k351), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk352() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k352), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k352), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk353() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k353), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k353), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk354() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k354), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k354), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk355() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k355), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k355), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk356() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k356), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k356), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk357() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k357), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k357), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk358() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k358), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k358), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk359() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k359), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k359), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk360() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k360), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k360), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk361() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k361), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k361), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk362() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k362), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k362), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk363() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k363), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k363), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk364() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k364), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k364), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk365() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k365), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k365), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk366() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k366), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k366), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk367() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k367), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k367), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk368() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k368), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k368), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk369() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k369), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k369), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk370() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k370), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k370), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk371() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k371), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k371), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk372() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k372), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k372), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk373() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k373), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k373), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk374() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k374), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k374), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk375() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k375), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k375), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk376() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k376), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k376), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk377() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k377), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k377), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk378() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k378), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k378), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk379() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k379), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k379), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk380() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k380), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k380), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk381() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k381), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k381), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk382() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k382), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k382), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk383() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k383), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k383), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk384() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k384), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k384), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk385() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k385), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k385), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk386() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k386), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k386), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk387() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k387), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k387), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk388() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k388), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k388), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk389() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k389), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k389), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk390() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k390), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k390), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk391() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k391), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k391), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk392() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k392), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k392), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk393() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k393), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k393), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk394() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k394), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k394), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk395() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k395), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k395), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk396() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k396), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k396), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk397() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k397), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k397), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk398() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k398), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k398), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk399() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k399), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k399), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk400() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k400), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k400), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk401() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k401), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k401), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk402() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k402), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k402), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk403() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k403), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k403), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk404() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k404), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k404), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk405() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k405), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k405), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk406() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k406), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k406), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk407() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k407), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k407), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk408() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k408), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k408), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk409() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k409), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k409), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk410() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k410), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k410), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk411() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k411), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k411), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk412() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k412), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k412), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk413() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k413), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k413), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk414() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k414), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k414), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk415() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k415), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k415), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk416() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k416), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k416), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk417() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k417), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k417), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk418() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k418), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k418), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk419() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k419), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k419), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk420() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k420), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k420), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk421() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k421), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k421), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk422() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k422), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k422), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk423() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k423), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k423), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk424() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k424), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k424), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk425() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k425), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k425), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk426() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k426), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k426), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk427() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k427), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k427), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk428() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k428), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k428), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk429() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k429), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k429), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk430() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k430), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k430), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk431() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k431), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k431), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk432() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k432), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k432), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk433() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k433), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k433), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk434() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k434), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k434), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk435() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k435), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k435), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk436() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k436), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k436), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk437() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k437), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k437), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk438() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k438), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k438), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk439() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k439), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k439), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk440() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k440), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k440), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk441() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k441), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k441), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk442() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k442), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k442), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk443() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k443), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k443), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk444() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k444), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k444), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk445() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k445), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k445), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk446() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k446), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k446), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk447() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k447), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k447), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk448() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k448), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k448), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk449() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k449), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k449), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk450() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k450), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k450), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk451() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k451), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k451), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk452() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k452), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k452), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk453() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k453), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k453), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk454() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k454), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k454), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk455() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k455), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k455), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk456() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k456), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k456), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk457() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k457), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k457), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk458() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k458), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k458), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk459() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k459), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k459), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk460() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k460), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k460), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk461() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k461), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k461), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk462() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k462), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k462), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk463() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k463), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k463), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk464() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k464), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k464), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk465() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k465), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k465), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk466() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k466), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k466), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk467() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k467), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k467), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk468() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k468), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k468), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk469() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k469), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k469), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk470() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k470), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k470), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk471() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k471), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k471), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk472() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k472), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k472), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk473() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k473), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k473), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk474() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k474), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k474), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk475() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k475), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k475), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk476() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k476), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k476), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk477() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k477), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k477), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk478() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k478), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k478), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk479() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k479), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k479), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk480() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k480), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k480), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk481() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k481), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k481), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk482() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k482), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k482), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk483() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k483), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k483), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk484() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k484), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k484), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk485() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k485), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k485), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk486() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k486), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k486), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk487() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k487), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k487), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk488() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k488), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k488), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk489() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k489), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k489), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk490() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k490), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k490), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk491() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k491), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k491), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk492() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k492), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k492), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk493() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k493), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k493), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk494() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k494), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k494), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk495() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k495), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k495), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk496() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k496), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k496), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk497() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k497), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k497), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk498() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k498), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k498), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}

func GetModeAmplitudeReImk499() []float64 {

	dm		:= Madd(&M, M0, 1.0, -1.0)

	vpRe	:= Dot(Cross(&M, psiIm_k499), dm)
	vpIm	:= Dot(Cross(&M, psiRe_k499), dm)

	valRe	:= ValueOf(vpRe)
	defer cuda.Recycle(valRe)

	valIm	:= ValueOf(vpIm)
	defer cuda.Recycle(valIm)

	amp := make([]float64, 3)

	ampRe	:= float64( cuda.Sum(valRe) )
	ampIm 	:= float64( cuda.Sum(valIm) )

	amp[0] = ampRe
	amp[1] = ampIm
	amp[2] = (ampRe * ampRe) + (ampIm * ampIm)

	return amp
}