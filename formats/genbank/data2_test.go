// Bulky test data #2.

package genbank

import "strings"

var want2 = &GenBank{
	Locus:       "AF165912                5485 bp    DNA     linear   PLN 29-JUL-1999",
	Definition:  "Arabidopsis thaliana CTP:phosphocholine cytidylyltransferase (CCT) gene, complete cds.",
	Accessions:  []string{"AF165912"},
	Version:     "AF165912.1",
	Keywords:    ".",
	Source:      "Arabidopsis thaliana (thale cress)",
	Organism:    "Arabidopsis thaliana",
	OrganismTax: "Eukaryota; Viridiplantae; Streptophyta; Embryophyta; Tracheophyta; Spermatophyta; Magnoliopsida; eudicotyledons; Gunneridae; Pentapetalae; rosids; malvids; Brassicales; Brassicaceae; Camelineae; Arabidopsis.",
	References: []map[string]string{
		{
			"":        "1  (bases 1 to 5485)",
			"AUTHORS": "Choi,Y.H., Choi,S.B. and Cho,S.H.",
			"TITLE":   "Structure of a CTP:Phosphocholine Cytidylyltransferase Gene from Arabidopsis thaliana",
			"JOURNAL": "Unpublished",
		},
		{
			"":        "2  (bases 1 to 5485)",
			"AUTHORS": "Choi,Y.H., Choi,S.B. and Cho,S.H.",
			"TITLE":   "Direct Submission",
			"JOURNAL": "Submitted (06-JUL-1999) Biology, Inha University, Yonghyon-Dong 253, Inchon 402-751, Korea",
		},
	},
	Features: []*Feature{
		{
			Type: "source",
			Fields: map[string]string{
				"":         "1..5485",
				"organism": "Arabidopsis thaliana",
				"mol_type": "genomic DNA",
				"db_xref":  "taxon:3702",
				"ecotype":  "Col-0",
			},
		},
		{
			Type: "gene",
			Fields: map[string]string{
				"":     "1..4637",
				"gene": "CCT",
			},
		},
		{
			Type: "regulatory",
			Fields: map[string]string{
				"":                 "1..1602",
				"regulatory_class": "promoter",
				"gene":             "CCT",
			},
		},
		{
			Type: "regulatory",
			Fields: map[string]string{
				"":                 "1554..1560",
				"regulatory_class": "TATA_box",
				"gene":             "CCT",
			},
		},
		{
			Type: "mRNA",
			Fields: map[string]string{
				"":        "join(1603..1891,2322..2438,2538..2633,2801..2843, 2918..3073,3167..3247,3874..3972,4082..4637)",
				"gene":    "CCT",
				"product": "CTP:phosphocholine cytidylyltransferase",
			},
		},
		{
			Type: "5'UTR",
			Fields: map[string]string{
				"":     "1603..1712",
				"gene": "CCT",
			},
		},
		{
			Type: "CDS",
			Fields: map[string]string{
				"":            "join(1713..1891,2322..2438,2538..2633,2801..2843, 2918..3073,3167..3247,3874..3972,4082..4309)",
				"gene":        "CCT",
				"EC_number":   "2.7.7.15",
				"codon_start": "1",
				"product":     "CTP:phosphocholine cytidylyltransferase",
				"protein_id":  "AAD45922.1",
				"translation": "MSNVIGDRTEDGLSTAAAASGSTAVQSSPPTDRPVRVYADGIYD" +
					"LFHFGHARSLEQAKLAFPNNTYLLVGCCNDETTHKYKGRTVMTAEERYESLRHCKWVD" +
					"EVIPDAPWVVNQEFLDKHQIDYVAHDSLPYADSSGAGKDVYEFVKKVGRFKETQRTEG" +
					"ISTSDIIMRIVKDYNQYVMRNLDRGYSREDLGVSFVKEKRLRVNMRLKKLQERVKEQQ" +
					"ERVGEKIQTVKMLRNEWVENADRWVAGFLEIFEEGCHKMGTAIVDSIQERLMRQKSAE" +
					"RLENGQDDDTDDQFYEEYFDHDMGSDDDEDEKFYDEEEVKEEETEKTVMTDAKDNK",
			},
		},
		{
			Type: "3'UTR",
			Fields: map[string]string{
				"":     "4310..4637",
				"gene": "CCT",
			},
		},
	},
	Origin: strings.ReplaceAll(origin2, "\n", ""),
}

// Taken from:
// https://www.ncbi.nlm.nih.gov/nucleotide/AF165912
const input2 = `LOCUS       AF165912                5485 bp    DNA     linear   PLN 29-JUL-1999
DEFINITION  Arabidopsis thaliana CTP:phosphocholine cytidylyltransferase (CCT)
            gene, complete cds.
ACCESSION   AF165912
VERSION     AF165912.1
KEYWORDS    .
SOURCE      Arabidopsis thaliana (thale cress)
  ORGANISM  Arabidopsis thaliana
            Eukaryota; Viridiplantae; Streptophyta; Embryophyta; Tracheophyta;
            Spermatophyta; Magnoliopsida; eudicotyledons; Gunneridae;
            Pentapetalae; rosids; malvids; Brassicales; Brassicaceae;
            Camelineae; Arabidopsis.
REFERENCE   1  (bases 1 to 5485)
  AUTHORS   Choi,Y.H., Choi,S.B. and Cho,S.H.
  TITLE     Structure of a CTP:Phosphocholine Cytidylyltransferase Gene from
            Arabidopsis thaliana
  JOURNAL   Unpublished
REFERENCE   2  (bases 1 to 5485)
  AUTHORS   Choi,Y.H., Choi,S.B. and Cho,S.H.
  TITLE     Direct Submission
  JOURNAL   Submitted (06-JUL-1999) Biology, Inha University, Yonghyon-Dong
            253, Inchon 402-751, Korea
FEATURES             Location/Qualifiers
     source          1..5485
                     /organism="Arabidopsis thaliana"
                     /mol_type="genomic DNA"
                     /db_xref="taxon:3702"
                     /ecotype="Col-0"
     gene            1..4637
                     /gene="CCT"
     regulatory      1..1602
                     /regulatory_class="promoter"
                     /gene="CCT"
     regulatory      1554..1560
                     /regulatory_class="TATA_box"
                     /gene="CCT"
     mRNA            join(1603..1891,2322..2438,2538..2633,2801..2843,
                     2918..3073,3167..3247,3874..3972,4082..4637)
                     /gene="CCT"
                     /product="CTP:phosphocholine cytidylyltransferase"
     5'UTR           1603..1712
                     /gene="CCT"
     CDS             join(1713..1891,2322..2438,2538..2633,2801..2843,
                     2918..3073,3167..3247,3874..3972,4082..4309)
                     /gene="CCT"
                     /EC_number="2.7.7.15"
                     /codon_start=1
                     /product="CTP:phosphocholine cytidylyltransferase"
                     /protein_id="AAD45922.1"
                     /translation="MSNVIGDRTEDGLSTAAAASGSTAVQSSPPTDRPVRVYADGIYD
                     LFHFGHARSLEQAKLAFPNNTYLLVGCCNDETTHKYKGRTVMTAEERYESLRHCKWVD
                     EVIPDAPWVVNQEFLDKHQIDYVAHDSLPYADSSGAGKDVYEFVKKVGRFKETQRTEG
                     ISTSDIIMRIVKDYNQYVMRNLDRGYSREDLGVSFVKEKRLRVNMRLKKLQERVKEQQ
                     ERVGEKIQTVKMLRNEWVENADRWVAGFLEIFEEGCHKMGTAIVDSIQERLMRQKSAE
                     RLENGQDDDTDDQFYEEYFDHDMGSDDDEDEKFYDEEEVKEEETEKTVMTDAKDNK"
     3'UTR           4310..4637
                     /gene="CCT"
ORIGIN      
        1 ccagaatggt tactatggac atccgccaac catacaagct atggtgaaat gctttatcta
       61 tctcattttt agtttcaaag cttttgttat aacacatgca aatccatatc cgtaaccaat
      121 atccaatcgc ttgacatagt ctgatgaagt ttttggtagt taagataaag ctcgagactg
      181 atatttcata tactggatga tttagggaaa cttgcattct attcatgaac gaatgagtca
      241 atacgagaca caaccaagca tgcaaggagc tgtgagttga tgttctatgc tatttaagta
      301 tttttcggga gatatatata tcttattgtt ctcctcctcc cgagtcaagt tgttctaaga
      361 aagaaggatc tatttcattt tgtggattgt ctagtttcag ggacagacgg ggtttagggg
      421 aagcgctatc cgtggctgct atgacatcga agaaactctg cacgacatgg tatgtaatct
      481 tgtgacgtta gtaaaaacgc tctaatgtta caaaagaaag aaagagaaaa cgaacccaat
      541 tcctcaaaat gttttctttt gacaatgtca acttcttcct tctcgggttt ttatcagttt
      601 gattgaagcg agactgcgaa attcctctgt ttacagtaga aaatgtgatc agccctattt
      661 ataaccgttg tatgttttcc ggtttttgtt tgtgcagaca atggggtcct cacagtttca
      721 gggatctgat tcgagccatc ctagtgatca ccgcttatcc aattaacaga acagaacaag
      781 ctcaagagtt gctactttca tatctttaaa atagtggaat gttttgtatg tacagaaata
      841 ggaaaggtct aaagtgtgga actggctttg aggtaaactc ttactctgat tggatttgct
      901 tgtatttata ccggaatcat aatagaaata tatgattaaa gtattcacat tctctaatct
      961 tcttttagac ttgtagttac cattcaaaag tcatggacaa ccttcgttaa ccttgagggc
     1021 cactgagaga caaccttgga cttggtcact ggctcactta tgcgtgcctc gccaaagtta
     1081 atctatcgac tgagattggt taatctggga acaaaaatta agagagaaag aagtagaact
     1141 aaaaagcaat attcagttat tcactctgtt gtatatagct caccaattaa attcaaacaa
     1201 ctaattcaaa acattatact atagcttttc attaaaaaat ttccaaaaca ttcatttaat
     1261 tatataaacc aaagaacagc ttttaagact taaaatattt cccaagtatc caacaagtca
     1321 atacagattt tttaagaaaa ctaaaccatt ttttcaactt tacaacaaaa acaccaactg
     1381 ttacaaaaaa actctcgaat ttcctatttc tccagcctta tgacaaagat atcagattaa
     1441 taaaatttag aattcattac tttttcttca tttttaaaat tatctacata ctatttattt
     1501 ctctccattt tattcagtga gaaaataaaa ttacaaatgc ctgaacacaa aaataaataa
     1561 aattagaata atcagtttcc tctagaggat attcttcgtc acaaaattaa aaaaaaaaag
     1621 agtaggagaa gagggaagcg actagcacct tttgtagttt tccgtttatt ttctgtataa
     1681 ggcgggtgat ttcggctcct tcatcggaaa ttatgagcaa cgttatcggc gatcgcactg
     1741 aagacggcct ttccaccgcc gctgcggcct ctggctctac ggctgtccag agttctcctc
     1801 ccactgatcg tcctgtccgc gtctacgccg atgggatcta cgatcttttc cactttggtc
     1861 atgctcgatc tctcgaacaa gccaaattag cgttagtatt tcttatctct tagagatgat
     1921 tgtcctgatt ttcatctcta attcgacttt tttttaccgt cgcgtgctca attttcgccg
     1981 ttccagtgtc atttttctct gatctgttga gtctggttca ttgtaagttg tacagttttt
     2041 gtttaggtcg agagacatat cttccttatt agatagtctc ggtgattgat tggtctgtat
     2101 tgattgaaat ctgtgatgtg caaggtcttg tcgcgtatga ttttagtgaa tccctttcta
     2161 aatgttgaaa tttgcaatag ctgatactgt ttctggatat atgttcttga cgaatgtttt
     2221 cgatttttta ttattttgag gaggtatgag agaaattgac ttctggtttc gtgttcttat
     2281 ggtgttgcta tgattgtgcc gtttcttaat cggccgagca ggtttccaaa caacacttac
     2341 cttcttgttg gatgttgcaa tgatgagact acccataagt acaagggaag gactgtaatg
     2401 actgcagaag agcgatatga atcacttcga cattgcaagt aattgttttc tcttatgttc
     2461 tgttgaatgt gttagtagaa aaacccatgg aagtggcagt gagtggaatt ttagaacacg
     2521 ttttttttat catgcaggtg ggtggatgaa gtcatccctg atgcaccatg ggtggtcaac
     2581 caggagtttc ttgacaagca ccagattgac tatgttgccc acgattctct tccgtaagaa
     2641 catgtgtctc ttgtgttagt ttttatttag ttttaaaaaa tggtgaaaac ttagttttgt
     2701 agtttttacc tttcacgacg tgcttgttgt tagtttagct cttttcttac aaatgatttt
     2761 agaactacaa taaccttctt tgtataattc tcatgcacag ctatgctgat tcaagcggag
     2821 ctggaaagga tgtctatgaa tttgtgagtc ggaagaattt tcatactcct gcttttgaca
     2881 ctttcatagt tctgttgtaa ctgagcatct gttgcaggtt aagaaagttg ggaggtttaa
     2941 ggaaacacag cgaactgaag gaatatcgac ctcggatata ataatgagaa tagtgaaaga
     3001 ttacaatcag tatgtcatgc gtaacttgga tagaggatac tcaagggaag atcttggagt
     3061 tagctttgtc aaggcatgtc atcattttct tatctctaca attttgtcct ttctcaaaaa
     3121 aaattcactt gtaagaatca actttggatt tgtcgatttg caacaggaaa agagacttag
     3181 agttaatatg aggctaaaga aactccagga gagggtcaaa gaacaacaag aaagagtggg
     3241 agaaaaggca tgtcttctct caacttcatt ttgcttaatt gatcattagt tcatcacaag
     3301 tccatcattt ggactgtatt gcattcaatc aaataaagct gttcatcata agttacaagg
     3361 agaaataact aaattttagg tcttgtctct gcctattcat tcacatctcc gcttgatctt
     3421 gtacctttga ctatttagcg actgtttgga aaccactctt aatgtgtcac gttttggagt
     3481 ctaacttgtc cttaatttga acctcgttca cttcttttag gactttaata ctctgtttgg
     3541 ttagtagcct ctaggcagaa aacatttgta tgtattgctt ttattttgtg tcttcttgtt
     3601 gtgattattg ggttatagaa ttgcatcaca aagtgatgct tgttaatccg ctgtagtagt
     3661 gccaggcgat atcatgttat ataatctcat ctcggtagta gcagccttat ctcgtgtatc
     3721 cgctgcgctt gaaacctcca tgcagtttca tgctttagct agtaatatga tatctgatga
     3781 gactaagttc atatgtgatt ctgaaaaagc tgattttgta gaagtttctt ataatgctcc
     3841 ttcctctgtt gttgttaaac ccggtttttc cagatccaaa ctgtaaaaat gctgcgcaac
     3901 gagtgggtag agaatgcaga tcgatgggtc gctggatttc ttgaaatatt tgaagaaggt
     3961 tgccataaga tggtaagttc aatcttgaag acacatacag tgcttcaaaa atctactaat
     4021 attcatgact atgttctgta taaccttgat taaacttgac aaatgcgtaa aatgttaaca
     4081 gggaactgca atcgtagaca gtatccaaga aaggttaatg agacaaaagt cggcagagag
     4141 gctggagaac ggtcaggatg atgacacaga cgaccagttc tatgaagaat acttcgatca
     4201 tgacatgggt agtgacgatg atgaagatga aaaattctac gacgaggaag aagtaaagga
     4261 agaagagaca gagaaaactg ttatgacgga tgctaaagac aacaagtaag aacaaatttg
     4321 gcttgcagaa acctcagatt agctctactt atggccactt ctactaaact cccttaagcc
     4381 tcgcactctc tctcgaaatt catctactta acatataata ccaatgttta gaaagagaga
     4441 gtgtgtgatg tgtttgtttg tgtgtgttga acaaacgaac gtgcgtggtt gtctttggtg
     4501 agttggtctc atctttgttg atttttgaat gcgcatgtat ttttttcttc tttttcatga
     4561 cgggcaaagt gttatgaagt acaatgcaat tgtctaaaac aggataagtc aatggttcgt
     4621 gtgtgccata aagtaaacat cgctgtgtac atcttccatg ttccaaactc aactcgtttt
     4681 cttcaaatat tgaaatacaa attggtcaaa agtcggttct tatttttttt ttaattcaca
     4741 tttttagttt gcagttttaa tagattacaa atcacatttt gtgctatttc caattccatg
     4801 agccggccaa gaatgtgagt aaaaggcaga taaagcaaag gatagccgat tgctttaaag
     4861 atgtctttgg taactagttc gaaattctct gtccactcga agactccaca actctcctct
     4921 caaatgtcag ctaatcaagt cctacacaac tatacaaaaa ggcaattaat tagtagaaaa
     4981 taaagattgg aggtttagct tctcccatac ataagtacct ttatgaatca ctaagctcag
     5041 ggtttatatg ataaccattg ctgatctgtg taaagagaag ttgatgaatt actacgtgag
     5101 tgttgttaac caactctctt tacatattag gaccgtgctt gtcaggccaa tggttttcac
     5161 ttcgaaaaat tgcttccgat atcaaactat gtgtacatta ttggtggact gtggacataa
     5221 cttaaacgca taattttatt gtgtaccttt aaaataaaca atagattaca catatatata
     5281 tggcaaatat ttgaacatta gatgtcaaga gaaaagtaaa acatgtcatg attacaccat
     5341 ctttgttatt atttagagtg attctcacta aatcttaggc ggttagcaac cgccatagtt
     5401 ttcaaaatct cattctatcg ggattaaatc tgtttttggt gactatatat aaacattggt
     5461 cgaattttta ggtaagtaaa atcag
//`

const origin2 = `ccagaatggttactatggacatccgccaaccatacaagctatggtgaaatgctttatcta
tctcatttttagtttcaaagcttttgttataacacatgcaaatccatatccgtaaccaat
atccaatcgcttgacatagtctgatgaagtttttggtagttaagataaagctcgagactg
atatttcatatactggatgatttagggaaacttgcattctattcatgaacgaatgagtca
atacgagacacaaccaagcatgcaaggagctgtgagttgatgttctatgctatttaagta
tttttcgggagatatatatatcttattgttctcctcctcccgagtcaagttgttctaaga
aagaaggatctatttcattttgtggattgtctagtttcagggacagacggggtttagggg
aagcgctatccgtggctgctatgacatcgaagaaactctgcacgacatggtatgtaatct
tgtgacgttagtaaaaacgctctaatgttacaaaagaaagaaagagaaaacgaacccaat
tcctcaaaatgttttcttttgacaatgtcaacttcttccttctcgggtttttatcagttt
gattgaagcgagactgcgaaattcctctgtttacagtagaaaatgtgatcagccctattt
ataaccgttgtatgttttccggtttttgtttgtgcagacaatggggtcctcacagtttca
gggatctgattcgagccatcctagtgatcaccgcttatccaattaacagaacagaacaag
ctcaagagttgctactttcatatctttaaaatagtggaatgttttgtatgtacagaaata
ggaaaggtctaaagtgtggaactggctttgaggtaaactcttactctgattggatttgct
tgtatttataccggaatcataatagaaatatatgattaaagtattcacattctctaatct
tcttttagacttgtagttaccattcaaaagtcatggacaaccttcgttaaccttgagggc
cactgagagacaaccttggacttggtcactggctcacttatgcgtgcctcgccaaagtta
atctatcgactgagattggttaatctgggaacaaaaattaagagagaaagaagtagaact
aaaaagcaatattcagttattcactctgttgtatatagctcaccaattaaattcaaacaa
ctaattcaaaacattatactatagcttttcattaaaaaatttccaaaacattcatttaat
tatataaaccaaagaacagcttttaagacttaaaatatttcccaagtatccaacaagtca
atacagattttttaagaaaactaaaccattttttcaactttacaacaaaaacaccaactg
ttacaaaaaaactctcgaatttcctatttctccagccttatgacaaagatatcagattaa
taaaatttagaattcattactttttcttcatttttaaaattatctacatactatttattt
ctctccattttattcagtgagaaaataaaattacaaatgcctgaacacaaaaataaataa
aattagaataatcagtttcctctagaggatattcttcgtcacaaaattaaaaaaaaaaag
agtaggagaagagggaagcgactagcaccttttgtagttttccgtttattttctgtataa
ggcgggtgatttcggctccttcatcggaaattatgagcaacgttatcggcgatcgcactg
aagacggcctttccaccgccgctgcggcctctggctctacggctgtccagagttctcctc
ccactgatcgtcctgtccgcgtctacgccgatgggatctacgatcttttccactttggtc
atgctcgatctctcgaacaagccaaattagcgttagtatttcttatctcttagagatgat
tgtcctgattttcatctctaattcgacttttttttaccgtcgcgtgctcaattttcgccg
ttccagtgtcatttttctctgatctgttgagtctggttcattgtaagttgtacagttttt
gtttaggtcgagagacatatcttccttattagatagtctcggtgattgattggtctgtat
tgattgaaatctgtgatgtgcaaggtcttgtcgcgtatgattttagtgaatccctttcta
aatgttgaaatttgcaatagctgatactgtttctggatatatgttcttgacgaatgtttt
cgattttttattattttgaggaggtatgagagaaattgacttctggtttcgtgttcttat
ggtgttgctatgattgtgccgtttcttaatcggccgagcaggtttccaaacaacacttac
cttcttgttggatgttgcaatgatgagactacccataagtacaagggaaggactgtaatg
actgcagaagagcgatatgaatcacttcgacattgcaagtaattgttttctcttatgttc
tgttgaatgtgttagtagaaaaacccatggaagtggcagtgagtggaattttagaacacg
ttttttttatcatgcaggtgggtggatgaagtcatccctgatgcaccatgggtggtcaac
caggagtttcttgacaagcaccagattgactatgttgcccacgattctcttccgtaagaa
catgtgtctcttgtgttagtttttatttagttttaaaaaatggtgaaaacttagttttgt
agtttttacctttcacgacgtgcttgttgttagtttagctcttttcttacaaatgatttt
agaactacaataaccttctttgtataattctcatgcacagctatgctgattcaagcggag
ctggaaaggatgtctatgaatttgtgagtcggaagaattttcatactcctgcttttgaca
ctttcatagttctgttgtaactgagcatctgttgcaggttaagaaagttgggaggtttaa
ggaaacacagcgaactgaaggaatatcgacctcggatataataatgagaatagtgaaaga
ttacaatcagtatgtcatgcgtaacttggatagaggatactcaagggaagatcttggagt
tagctttgtcaaggcatgtcatcattttcttatctctacaattttgtcctttctcaaaaa
aaattcacttgtaagaatcaactttggatttgtcgatttgcaacaggaaaagagacttag
agttaatatgaggctaaagaaactccaggagagggtcaaagaacaacaagaaagagtggg
agaaaaggcatgtcttctctcaacttcattttgcttaattgatcattagttcatcacaag
tccatcatttggactgtattgcattcaatcaaataaagctgttcatcataagttacaagg
agaaataactaaattttaggtcttgtctctgcctattcattcacatctccgcttgatctt
gtacctttgactatttagcgactgtttggaaaccactcttaatgtgtcacgttttggagt
ctaacttgtccttaatttgaacctcgttcacttcttttaggactttaatactctgtttgg
ttagtagcctctaggcagaaaacatttgtatgtattgcttttattttgtgtcttcttgtt
gtgattattgggttatagaattgcatcacaaagtgatgcttgttaatccgctgtagtagt
gccaggcgatatcatgttatataatctcatctcggtagtagcagccttatctcgtgtatc
cgctgcgcttgaaacctccatgcagtttcatgctttagctagtaatatgatatctgatga
gactaagttcatatgtgattctgaaaaagctgattttgtagaagtttcttataatgctcc
ttcctctgttgttgttaaacccggtttttccagatccaaactgtaaaaatgctgcgcaac
gagtgggtagagaatgcagatcgatgggtcgctggatttcttgaaatatttgaagaaggt
tgccataagatggtaagttcaatcttgaagacacatacagtgcttcaaaaatctactaat
attcatgactatgttctgtataaccttgattaaacttgacaaatgcgtaaaatgttaaca
gggaactgcaatcgtagacagtatccaagaaaggttaatgagacaaaagtcggcagagag
gctggagaacggtcaggatgatgacacagacgaccagttctatgaagaatacttcgatca
tgacatgggtagtgacgatgatgaagatgaaaaattctacgacgaggaagaagtaaagga
agaagagacagagaaaactgttatgacggatgctaaagacaacaagtaagaacaaatttg
gcttgcagaaacctcagattagctctacttatggccacttctactaaactcccttaagcc
tcgcactctctctcgaaattcatctacttaacatataataccaatgtttagaaagagaga
gtgtgtgatgtgtttgtttgtgtgtgttgaacaaacgaacgtgcgtggttgtctttggtg
agttggtctcatctttgttgatttttgaatgcgcatgtatttttttcttctttttcatga
cgggcaaagtgttatgaagtacaatgcaattgtctaaaacaggataagtcaatggttcgt
gtgtgccataaagtaaacatcgctgtgtacatcttccatgttccaaactcaactcgtttt
cttcaaatattgaaatacaaattggtcaaaagtcggttcttattttttttttaattcaca
tttttagtttgcagttttaatagattacaaatcacattttgtgctatttccaattccatg
agccggccaagaatgtgagtaaaaggcagataaagcaaaggatagccgattgctttaaag
atgtctttggtaactagttcgaaattctctgtccactcgaagactccacaactctcctct
caaatgtcagctaatcaagtcctacacaactatacaaaaaggcaattaattagtagaaaa
taaagattggaggtttagcttctcccatacataagtacctttatgaatcactaagctcag
ggtttatatgataaccattgctgatctgtgtaaagagaagttgatgaattactacgtgag
tgttgttaaccaactctctttacatattaggaccgtgcttgtcaggccaatggttttcac
ttcgaaaaattgcttccgatatcaaactatgtgtacattattggtggactgtggacataa
cttaaacgcataattttattgtgtacctttaaaataaacaatagattacacatatatata
tggcaaatatttgaacattagatgtcaagagaaaagtaaaacatgtcatgattacaccat
ctttgttattatttagagtgattctcactaaatcttaggcggttagcaaccgccatagtt
ttcaaaatctcattctatcgggattaaatctgtttttggtgactatatataaacattggt
cgaatttttaggtaagtaaaatcag`
