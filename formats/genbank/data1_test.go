// Bulky test data #1.

package genbank

import "strings"

var want1 = &GenBank{
	Locus:       "SCU49845     5028 bp    DNA             PLN       21-JUN-1999",
	Definition:  "Saccharomyces cerevisiae TCP1-beta gene, partial cds, and Axl2p (AXL2) and Rev7p (REV7) genes, complete cds.",
	Accessions:  []string{"U49845"},
	Version:     "U49845.1  GI:1293613",
	Keywords:    ".",
	Source:      "Saccharomyces cerevisiae (baker's yeast)",
	Organism:    "Saccharomyces cerevisiae",
	OrganismTax: "Eukaryota; Fungi; Ascomycota; Saccharomycotina; Saccharomycetes; Saccharomycetales; Saccharomycetaceae; Saccharomyces.",
	References: []map[string]string{
		{
			"":        "1  (bases 1 to 5028)",
			"AUTHORS": "Torpey,L.E., Gibbs,P.E., Nelson,J. and Lawrence,C.W.",
			"TITLE":   "Cloning and sequence of REV7, a gene whose function is required for DNA damage-induced mutagenesis in Saccharomyces cerevisiae",
			"JOURNAL": "Yeast 10 (11), 1503-1509 (1994)",
			"PUBMED":  "7871890",
		},
		{
			"":        "2  (bases 1 to 5028)",
			"AUTHORS": "Roemer,T., Madden,K., Chang,J. and Snyder,M.",
			"TITLE":   "Selection of axial growth sites in yeast requires Axl2p, a novel plasma membrane glycoprotein",
			"JOURNAL": "Genes Dev. 10 (7), 777-793 (1996)",
			"PUBMED":  "8846915",
		},
		{
			"":        "3  (bases 1 to 5028)",
			"AUTHORS": "Roemer,T.",
			"TITLE":   "Direct Submission",
			"JOURNAL": "Submitted (22-FEB-1996) Terry Roemer, Biology, Yale University, New Haven, CT, USA",
		},
	},
	Features: []*Feature{
		{
			Type: "source",
			Fields: map[string]string{
				"":           "1..5028",
				"organism":   "Saccharomyces cerevisiae",
				"db_xref":    "taxon:4932",
				"chromosome": "IX",
				"map":        "9",
			},
		},
		{
			Type: "CDS",
			Fields: map[string]string{
				"":            "<1..206",
				"codon_start": "3",
				"product":     "TCP1-beta",
				"protein_id":  "AAA98665.1",
				"db_xref":     "GI:1293614",
				"translation": "SSIYNGISTSGLDLNNGTIADMRQLGIVESYKLKRAVVSSASEAAEVLLRVDNIIRARPRTANRQHM",
			},
		},
		{
			Type: "gene",
			Fields: map[string]string{
				"":     "687..3158",
				"gene": "AXL2",
			},
		},
		{
			Type: "CDS",
			Fields: map[string]string{
				"":            "687..3158",
				"gene":        "AXL2",
				"note":        "plasma membrane glycoprotein",
				"codon_start": "1",
				"function":    "required for axial budding pattern of S. cerevisiae",
				"product":     "Axl2p",
				"protein_id":  "AAA98666.1",
				"db_xref":     "GI:1293615",
				"translation": "MTQLQISLLLTATISLLHLVVATPYEAYPIGKQYPPVARVNESF" +
					"TFQISNDTYKSSVDKTAQITYNCFDLPSWLSFDSSSRTFSGEPSSDLLSDANTTLYFN" +
					"VILEGTDSADSTSLNNTYQFVVTNRPSISLSSDFNLLALLKNYGYTNGKNALKLDPNE" +
					"VFNVTFDRSMFTNEESIVSYYGRSQLYNAPLPNWLFFDSGELKFTGTAPVINSAIAPE" +
					"TSYSFVIIATDIEGFSAVEVEFELVIGAHQLTTSIQNSLIINVTDTGNVSYDLPLNYV" +
					"YLDDDPISSDKLGSINLLDAPDWVALDNATISGSVPDELLGKNSNPANFSVSIYDTYG" +
					"DVIYFNFEVVSTTDLFAISSLPNINATRGEWFSYYFLPSQFTDYVNTNVSLEFTNSSQ" +
					"DHDWVKFQSSNLTLAGEVPKNFDKLSLGLKANQGSQSQELYFNIIGMDSKITHSNHSA" +
					"NATSTRSSHHSTSTSSYTSSTYTAKISSTSAAATSSAPAALPAANKTSSHNKKAVAIA" +
					"CGVAIPLGVILVALICFLIFWRRRRENPDDENLPHAISGPDLNNPANKPNQENATPLN" +
					"NPFDDDASSYDDTSIARRLAALNTLKLDNHSATESDISSVDEKRDSLSGMNTYNDQFQ" +
					"SQSKEELLAKPPVQPPESPFFDPQNRSSSVYMDSEPAVNKSWRYTGNLSPVSDIVRDS" +
					"YGSQKTVDTEKLFDLEAPEKEKRTSRDVTMSSLDPWNSNISPSPVRKSVTPSPYNVTK" +
					"HRNRHLQNIQDSQSGKNGITPTTMSTSSSDDFVPVKDGENFCWVHSMEPDRRPSKKRL" +
					"VDFSNKSNVNVGQVKDIHGRIPEML",
			},
		},
		{
			Type: "gene",
			Fields: map[string]string{
				"":     "complement(3300..4037)",
				"gene": "REV7",
			},
		},
		{
			Type: "CDS",
			Fields: map[string]string{
				"":            "complement(3300..4037)",
				"gene":        "REV7",
				"codon_start": "1",
				"product":     "Rev7p",
				"protein_id":  "AAA98667.1",
				"db_xref":     "GI:1293616",
				"translation": "MNRWVEKWLRVYLKCYINLILFYRNVYPPQSFDYTTYQSFNLPQ" +
					"FVPINRHPALIDYIEELILDVLSKLTHVYRFSICIINKKNDLCIEKYVLDFSELQHVD" +
					"KDDQIITETEVFDEFRSSLNSLIMHLEKLPKVNDDTITFEAVINAIELELGHKLDRNR" +
					"RVDSLEEKAEIERDSNWVKCQEDENLPDNNGFQPPKIKLTSLVGSDVGPLIIHQFSEK" +
					"LISGDDKILNGVYSQYEEGESIFGSLF",
			},
		},
	},
	Origin: strings.ReplaceAll(origin1, "\n", ""),
}

// Taken from:
// https://www.ncbi.nlm.nih.gov/genbank/samplerecord
const input1 = `LOCUS       SCU49845     5028 bp    DNA             PLN       21-JUN-1999
DEFINITION  Saccharomyces cerevisiae TCP1-beta gene, partial cds, and Axl2p
            (AXL2) and Rev7p (REV7) genes, complete cds.
ACCESSION   U49845
VERSION     U49845.1  GI:1293613
KEYWORDS    .
SOURCE      Saccharomyces cerevisiae (baker's yeast)
  ORGANISM  Saccharomyces cerevisiae
            Eukaryota; Fungi; Ascomycota; Saccharomycotina; Saccharomycetes;
            Saccharomycetales; Saccharomycetaceae; Saccharomyces.
REFERENCE   1  (bases 1 to 5028)
  AUTHORS   Torpey,L.E., Gibbs,P.E., Nelson,J. and Lawrence,C.W.
  TITLE     Cloning and sequence of REV7, a gene whose function is required for
            DNA damage-induced mutagenesis in Saccharomyces cerevisiae
  JOURNAL   Yeast 10 (11), 1503-1509 (1994)
  PUBMED    7871890
REFERENCE   2  (bases 1 to 5028)
  AUTHORS   Roemer,T., Madden,K., Chang,J. and Snyder,M.
  TITLE     Selection of axial growth sites in yeast requires Axl2p, a novel
            plasma membrane glycoprotein
  JOURNAL   Genes Dev. 10 (7), 777-793 (1996)
  PUBMED    8846915
REFERENCE   3  (bases 1 to 5028)
  AUTHORS   Roemer,T.
  TITLE     Direct Submission
  JOURNAL   Submitted (22-FEB-1996) Terry Roemer, Biology, Yale University, New
            Haven, CT, USA
FEATURES             Location/Qualifiers
     source          1..5028
                     /organism="Saccharomyces cerevisiae"
                     /db_xref="taxon:4932"
                     /chromosome="IX"
                     /map="9"
     CDS             <1..206
                     /codon_start=3
                     /product="TCP1-beta"
                     /protein_id="AAA98665.1"
                     /db_xref="GI:1293614"
                     /translation="SSIYNGISTSGLDLNNGTIADMRQLGIVESYKLKRAVVSSASEA
                     AEVLLRVDNIIRARPRTANRQHM"
     gene            687..3158
                     /gene="AXL2"
     CDS             687..3158
                     /gene="AXL2"
                     /note="plasma membrane glycoprotein"
                     /codon_start=1
                     /function="required for axial budding pattern of S.
                     cerevisiae"
                     /product="Axl2p"
                     /protein_id="AAA98666.1"
                     /db_xref="GI:1293615"
                     /translation="MTQLQISLLLTATISLLHLVVATPYEAYPIGKQYPPVARVNESF
                     TFQISNDTYKSSVDKTAQITYNCFDLPSWLSFDSSSRTFSGEPSSDLLSDANTTLYFN
                     VILEGTDSADSTSLNNTYQFVVTNRPSISLSSDFNLLALLKNYGYTNGKNALKLDPNE
                     VFNVTFDRSMFTNEESIVSYYGRSQLYNAPLPNWLFFDSGELKFTGTAPVINSAIAPE
                     TSYSFVIIATDIEGFSAVEVEFELVIGAHQLTTSIQNSLIINVTDTGNVSYDLPLNYV
                     YLDDDPISSDKLGSINLLDAPDWVALDNATISGSVPDELLGKNSNPANFSVSIYDTYG
                     DVIYFNFEVVSTTDLFAISSLPNINATRGEWFSYYFLPSQFTDYVNTNVSLEFTNSSQ
                     DHDWVKFQSSNLTLAGEVPKNFDKLSLGLKANQGSQSQELYFNIIGMDSKITHSNHSA
                     NATSTRSSHHSTSTSSYTSSTYTAKISSTSAAATSSAPAALPAANKTSSHNKKAVAIA
                     CGVAIPLGVILVALICFLIFWRRRRENPDDENLPHAISGPDLNNPANKPNQENATPLN
                     NPFDDDASSYDDTSIARRLAALNTLKLDNHSATESDISSVDEKRDSLSGMNTYNDQFQ
                     SQSKEELLAKPPVQPPESPFFDPQNRSSSVYMDSEPAVNKSWRYTGNLSPVSDIVRDS
                     YGSQKTVDTEKLFDLEAPEKEKRTSRDVTMSSLDPWNSNISPSPVRKSVTPSPYNVTK
                     HRNRHLQNIQDSQSGKNGITPTTMSTSSSDDFVPVKDGENFCWVHSMEPDRRPSKKRL
                     VDFSNKSNVNVGQVKDIHGRIPEML"
     gene            complement(3300..4037)
                     /gene="REV7"
     CDS             complement(3300..4037)
                     /gene="REV7"
                     /codon_start=1
                     /product="Rev7p"
                     /protein_id="AAA98667.1"
                     /db_xref="GI:1293616"
                     /translation="MNRWVEKWLRVYLKCYINLILFYRNVYPPQSFDYTTYQSFNLPQ
                     FVPINRHPALIDYIEELILDVLSKLTHVYRFSICIINKKNDLCIEKYVLDFSELQHVD
                     KDDQIITETEVFDEFRSSLNSLIMHLEKLPKVNDDTITFEAVINAIELELGHKLDRNR
                     RVDSLEEKAEIERDSNWVKCQEDENLPDNNGFQPPKIKLTSLVGSDVGPLIIHQFSEK
                     LISGDDKILNGVYSQYEEGESIFGSLF"
ORIGIN
        1 gatcctccat atacaacggt atctccacct caggtttaga tctcaacaac ggaaccattg
       61 ccgacatgag acagttaggt atcgtcgaga gttacaagct aaaacgagca gtagtcagct
      121 ctgcatctga agccgctgaa gttctactaa gggtggataa catcatccgt gcaagaccaa
      181 gaaccgccaa tagacaacat atgtaacata tttaggatat acctcgaaaa taataaaccg
      241 ccacactgtc attattataa ttagaaacag aacgcaaaaa ttatccacta tataattcaa
      301 agacgcgaaa aaaaaagaac aacgcgtcat agaacttttg gcaattcgcg tcacaaataa
      361 attttggcaa cttatgtttc ctcttcgagc agtactcgag ccctgtctca agaatgtaat
      421 aatacccatc gtaggtatgg ttaaagatag catctccaca acctcaaagc tccttgccga
      481 gagtcgccct cctttgtcga gtaattttca cttttcatat gagaacttat tttcttattc
      541 tttactctca catcctgtag tgattgacac tgcaacagcc accatcacta gaagaacaga
      601 acaattactt aatagaaaaa ttatatcttc ctcgaaacga tttcctgctt ccaacatcta
      661 cgtatatcaa gaagcattca cttaccatga cacagcttca gatttcatta ttgctgacag
      721 ctactatatc actactccat ctagtagtgg ccacgcccta tgaggcatat cctatcggaa
      781 aacaataccc cccagtggca agagtcaatg aatcgtttac atttcaaatt tccaatgata
      841 cctataaatc gtctgtagac aagacagctc aaataacata caattgcttc gacttaccga
      901 gctggctttc gtttgactct agttctagaa cgttctcagg tgaaccttct tctgacttac
      961 tatctgatgc gaacaccacg ttgtatttca atgtaatact cgagggtacg gactctgccg
     1021 acagcacgtc tttgaacaat acataccaat ttgttgttac aaaccgtcca tccatctcgc
     1081 tatcgtcaga tttcaatcta ttggcgttgt taaaaaacta tggttatact aacggcaaaa
     1141 acgctctgaa actagatcct aatgaagtct tcaacgtgac ttttgaccgt tcaatgttca
     1201 ctaacgaaga atccattgtg tcgtattacg gacgttctca gttgtataat gcgccgttac
     1261 ccaattggct gttcttcgat tctggcgagt tgaagtttac tgggacggca ccggtgataa
     1321 actcggcgat tgctccagaa acaagctaca gttttgtcat catcgctaca gacattgaag
     1381 gattttctgc cgttgaggta gaattcgaat tagtcatcgg ggctcaccag ttaactacct
     1441 ctattcaaaa tagtttgata atcaacgtta ctgacacagg taacgtttca tatgacttac
     1501 ctctaaacta tgtttatctc gatgacgatc ctatttcttc tgataaattg ggttctataa
     1561 acttattgga tgctccagac tgggtggcat tagataatgc taccatttcc gggtctgtcc
     1621 cagatgaatt actcggtaag aactccaatc ctgccaattt ttctgtgtcc atttatgata
     1681 cttatggtga tgtgatttat ttcaacttcg aagttgtctc cacaacggat ttgtttgcca
     1741 ttagttctct tcccaatatt aacgctacaa ggggtgaatg gttctcctac tattttttgc
     1801 cttctcagtt tacagactac gtgaatacaa acgtttcatt agagtttact aattcaagcc
     1861 aagaccatga ctgggtgaaa ttccaatcat ctaatttaac attagctgga gaagtgccca
     1921 agaatttcga caagctttca ttaggtttga aagcgaacca aggttcacaa tctcaagagc
     1981 tatattttaa catcattggc atggattcaa agataactca ctcaaaccac agtgcgaatg
     2041 caacgtccac aagaagttct caccactcca cctcaacaag ttcttacaca tcttctactt
     2101 acactgcaaa aatttcttct acctccgctg ctgctacttc ttctgctcca gcagcgctgc
     2161 cagcagccaa taaaacttca tctcacaata aaaaagcagt agcaattgcg tgcggtgttg
     2221 ctatcccatt aggcgttatc ctagtagctc tcatttgctt cctaatattc tggagacgca
     2281 gaagggaaaa tccagacgat gaaaacttac cgcatgctat tagtggacct gatttgaata
     2341 atcctgcaaa taaaccaaat caagaaaacg ctacaccttt gaacaacccc tttgatgatg
     2401 atgcttcctc gtacgatgat acttcaatag caagaagatt ggctgctttg aacactttga
     2461 aattggataa ccactctgcc actgaatctg atatttccag cgtggatgaa aagagagatt
     2521 ctctatcagg tatgaataca tacaatgatc agttccaatc ccaaagtaaa gaagaattat
     2581 tagcaaaacc cccagtacag cctccagaga gcccgttctt tgacccacag aataggtctt
     2641 cttctgtgta tatggatagt gaaccagcag taaataaatc ctggcgatat actggcaacc
     2701 tgtcaccagt ctctgatatt gtcagagaca gttacggatc acaaaaaact gttgatacag
     2761 aaaaactttt cgatttagaa gcaccagaga aggaaaaacg tacgtcaagg gatgtcacta
     2821 tgtcttcact ggacccttgg aacagcaata ttagcccttc tcccgtaaga aaatcagtaa
     2881 caccatcacc atataacgta acgaagcatc gtaaccgcca cttacaaaat attcaagact
     2941 ctcaaagcgg taaaaacgga atcactccca caacaatgtc aacttcatct tctgacgatt
     3001 ttgttccggt taaagatggt gaaaattttt gctgggtcca tagcatggaa ccagacagaa
     3061 gaccaagtaa gaaaaggtta gtagattttt caaataagag taatgtcaat gttggtcaag
     3121 ttaaggacat tcacggacgc atcccagaaa tgctgtgatt atacgcaacg atattttgct
     3181 taattttatt ttcctgtttt attttttatt agtggtttac agatacccta tattttattt
     3241 agtttttata cttagagaca tttaatttta attccattct tcaaatttca tttttgcact
     3301 taaaacaaag atccaaaaat gctctcgccc tcttcatatt gagaatacac tccattcaaa
     3361 attttgtcgt caccgctgat taatttttca ctaaactgat gaataatcaa aggccccacg
     3421 tcagaaccga ctaaagaagt gagttttatt ttaggaggtt gaaaaccatt attgtctggt
     3481 aaattttcat cttcttgaca tttaacccag tttgaatccc tttcaatttc tgctttttcc
     3541 tccaaactat cgaccctcct gtttctgtcc aacttatgtc ctagttccaa ttcgatcgca
     3601 ttaataactg cttcaaatgt tattgtgtca tcgttgactt taggtaattt ctccaaatgc
     3661 ataatcaaac tatttaagga agatcggaat tcgtcgaaca cttcagtttc cgtaatgatc
     3721 tgatcgtctt tatccacatg ttgtaattca ctaaaatcta aaacgtattt ttcaatgcat
     3781 aaatcgttct ttttattaat aatgcagatg gaaaatctgt aaacgtgcgt taatttagaa
     3841 agaacatcca gtataagttc ttctatatag tcaattaaag caggatgcct attaatggga
     3901 acgaactgcg gcaagttgaa tgactggtaa gtagtgtagt cgaatgactg aggtgggtat
     3961 acatttctat aaaataaaat caaattaatg tagcatttta agtataccct cagccacttc
     4021 tctacccatc tattcataaa gctgacgcaa cgattactat tttttttttc ttcttggatc
     4081 tcagtcgtcg caaaaacgta taccttcttt ttccgacctt ttttttagct ttctggaaaa
     4141 gtttatatta gttaaacagg gtctagtctt agtgtgaaag ctagtggttt cgattgactg
     4201 atattaagaa agtggaaatt aaattagtag tgtagacgta tatgcatatg tatttctcgc
     4261 ctgtttatgt ttctacgtac ttttgattta tagcaagggg aaaagaaata catactattt
     4321 tttggtaaag gtgaaagcat aatgtaaaag ctagaataaa atggacgaaa taaagagagg
     4381 cttagttcat cttttttcca aaaagcaccc aatgataata actaaaatga aaaggatttg
     4441 ccatctgtca gcaacatcag ttgtgtgagc aataataaaa tcatcacctc cgttgccttt
     4501 agcgcgtttg tcgtttgtat cttccgtaat tttagtctta tcaatgggaa tcataaattt
     4561 tccaatgaat tagcaatttc gtccaattct ttttgagctt cttcatattt gctttggaat
     4621 tcttcgcact tcttttccca ttcatctctt tcttcttcca aagcaacgat ccttctaccc
     4681 atttgctcag agttcaaatc ggcctctttc agtttatcca ttgcttcctt cagtttggct
     4741 tcactgtctt ctagctgttg ttctagatcc tggtttttct tggtgtagtt ctcattatta
     4801 gatctcaagt tattggagtc ttcagccaat tgctttgtat cagacaattg actctctaac
     4861 ttctccactt cactgtcgag ttgctcgttt ttagcggaca aagatttaat ctcgttttct
     4921 ttttcagtgt tagattgctc taattctttg agctgttctc tcagctcctc atatttttct
     4981 tgccatgact cagattctaa ttttaagcta ttcaatttct ctttgatc
//

`

const origin1 = `gatcctccatatacaacggtatctccacctcaggtttagatctcaacaacggaaccattg
ccgacatgagacagttaggtatcgtcgagagttacaagctaaaacgagcagtagtcagct
ctgcatctgaagccgctgaagttctactaagggtggataacatcatccgtgcaagaccaa
gaaccgccaatagacaacatatgtaacatatttaggatatacctcgaaaataataaaccg
ccacactgtcattattataattagaaacagaacgcaaaaattatccactatataattcaa
agacgcgaaaaaaaaagaacaacgcgtcatagaacttttggcaattcgcgtcacaaataa
attttggcaacttatgtttcctcttcgagcagtactcgagccctgtctcaagaatgtaat
aatacccatcgtaggtatggttaaagatagcatctccacaacctcaaagctccttgccga
gagtcgccctcctttgtcgagtaattttcacttttcatatgagaacttattttcttattc
tttactctcacatcctgtagtgattgacactgcaacagccaccatcactagaagaacaga
acaattacttaatagaaaaattatatcttcctcgaaacgatttcctgcttccaacatcta
cgtatatcaagaagcattcacttaccatgacacagcttcagatttcattattgctgacag
ctactatatcactactccatctagtagtggccacgccctatgaggcatatcctatcggaa
aacaataccccccagtggcaagagtcaatgaatcgtttacatttcaaatttccaatgata
cctataaatcgtctgtagacaagacagctcaaataacatacaattgcttcgacttaccga
gctggctttcgtttgactctagttctagaacgttctcaggtgaaccttcttctgacttac
tatctgatgcgaacaccacgttgtatttcaatgtaatactcgagggtacggactctgccg
acagcacgtctttgaacaatacataccaatttgttgttacaaaccgtccatccatctcgc
tatcgtcagatttcaatctattggcgttgttaaaaaactatggttatactaacggcaaaa
acgctctgaaactagatcctaatgaagtcttcaacgtgacttttgaccgttcaatgttca
ctaacgaagaatccattgtgtcgtattacggacgttctcagttgtataatgcgccgttac
ccaattggctgttcttcgattctggcgagttgaagtttactgggacggcaccggtgataa
actcggcgattgctccagaaacaagctacagttttgtcatcatcgctacagacattgaag
gattttctgccgttgaggtagaattcgaattagtcatcggggctcaccagttaactacct
ctattcaaaatagtttgataatcaacgttactgacacaggtaacgtttcatatgacttac
ctctaaactatgtttatctcgatgacgatcctatttcttctgataaattgggttctataa
acttattggatgctccagactgggtggcattagataatgctaccatttccgggtctgtcc
cagatgaattactcggtaagaactccaatcctgccaatttttctgtgtccatttatgata
cttatggtgatgtgatttatttcaacttcgaagttgtctccacaacggatttgtttgcca
ttagttctcttcccaatattaacgctacaaggggtgaatggttctcctactattttttgc
cttctcagtttacagactacgtgaatacaaacgtttcattagagtttactaattcaagcc
aagaccatgactgggtgaaattccaatcatctaatttaacattagctggagaagtgccca
agaatttcgacaagctttcattaggtttgaaagcgaaccaaggttcacaatctcaagagc
tatattttaacatcattggcatggattcaaagataactcactcaaaccacagtgcgaatg
caacgtccacaagaagttctcaccactccacctcaacaagttcttacacatcttctactt
acactgcaaaaatttcttctacctccgctgctgctacttcttctgctccagcagcgctgc
cagcagccaataaaacttcatctcacaataaaaaagcagtagcaattgcgtgcggtgttg
ctatcccattaggcgttatcctagtagctctcatttgcttcctaatattctggagacgca
gaagggaaaatccagacgatgaaaacttaccgcatgctattagtggacctgatttgaata
atcctgcaaataaaccaaatcaagaaaacgctacacctttgaacaacccctttgatgatg
atgcttcctcgtacgatgatacttcaatagcaagaagattggctgctttgaacactttga
aattggataaccactctgccactgaatctgatatttccagcgtggatgaaaagagagatt
ctctatcaggtatgaatacatacaatgatcagttccaatcccaaagtaaagaagaattat
tagcaaaacccccagtacagcctccagagagcccgttctttgacccacagaataggtctt
cttctgtgtatatggatagtgaaccagcagtaaataaatcctggcgatatactggcaacc
tgtcaccagtctctgatattgtcagagacagttacggatcacaaaaaactgttgatacag
aaaaacttttcgatttagaagcaccagagaaggaaaaacgtacgtcaagggatgtcacta
tgtcttcactggacccttggaacagcaatattagcccttctcccgtaagaaaatcagtaa
caccatcaccatataacgtaacgaagcatcgtaaccgccacttacaaaatattcaagact
ctcaaagcggtaaaaacggaatcactcccacaacaatgtcaacttcatcttctgacgatt
ttgttccggttaaagatggtgaaaatttttgctgggtccatagcatggaaccagacagaa
gaccaagtaagaaaaggttagtagatttttcaaataagagtaatgtcaatgttggtcaag
ttaaggacattcacggacgcatcccagaaatgctgtgattatacgcaacgatattttgct
taattttattttcctgttttattttttattagtggtttacagataccctatattttattt
agtttttatacttagagacatttaattttaattccattcttcaaatttcatttttgcact
taaaacaaagatccaaaaatgctctcgccctcttcatattgagaatacactccattcaaa
attttgtcgtcaccgctgattaatttttcactaaactgatgaataatcaaaggccccacg
tcagaaccgactaaagaagtgagttttattttaggaggttgaaaaccattattgtctggt
aaattttcatcttcttgacatttaacccagtttgaatccctttcaatttctgctttttcc
tccaaactatcgaccctcctgtttctgtccaacttatgtcctagttccaattcgatcgca
ttaataactgcttcaaatgttattgtgtcatcgttgactttaggtaatttctccaaatgc
ataatcaaactatttaaggaagatcggaattcgtcgaacacttcagtttccgtaatgatc
tgatcgtctttatccacatgttgtaattcactaaaatctaaaacgtatttttcaatgcat
aaatcgttctttttattaataatgcagatggaaaatctgtaaacgtgcgttaatttagaa
agaacatccagtataagttcttctatatagtcaattaaagcaggatgcctattaatggga
acgaactgcggcaagttgaatgactggtaagtagtgtagtcgaatgactgaggtgggtat
acatttctataaaataaaatcaaattaatgtagcattttaagtataccctcagccacttc
tctacccatctattcataaagctgacgcaacgattactattttttttttcttcttggatc
tcagtcgtcgcaaaaacgtataccttctttttccgaccttttttttagctttctggaaaa
gtttatattagttaaacagggtctagtcttagtgtgaaagctagtggtttcgattgactg
atattaagaaagtggaaattaaattagtagtgtagacgtatatgcatatgtatttctcgc
ctgtttatgtttctacgtacttttgatttatagcaaggggaaaagaaatacatactattt
tttggtaaaggtgaaagcataatgtaaaagctagaataaaatggacgaaataaagagagg
cttagttcatcttttttccaaaaagcacccaatgataataactaaaatgaaaaggatttg
ccatctgtcagcaacatcagttgtgtgagcaataataaaatcatcacctccgttgccttt
agcgcgtttgtcgtttgtatcttccgtaattttagtcttatcaatgggaatcataaattt
tccaatgaattagcaatttcgtccaattctttttgagcttcttcatatttgctttggaat
tcttcgcacttcttttcccattcatctctttcttcttccaaagcaacgatccttctaccc
atttgctcagagttcaaatcggcctctttcagtttatccattgcttccttcagtttggct
tcactgtcttctagctgttgttctagatcctggtttttcttggtgtagttctcattatta
gatctcaagttattggagtcttcagccaattgctttgtatcagacaattgactctctaac
ttctccacttcactgtcgagttgctcgtttttagcggacaaagatttaatctcgttttct
ttttcagtgttagattgctctaattctttgagctgttctctcagctcctcatatttttct
tgccatgactcagattctaattttaagctattcaatttctctttgatc`
