// Bulky test data #3.

package genbank

var want3 = &GenBank{
	Locus:       "AF090832                5086 bp    DNA     linear   INV 04-AUG-1999",
	Definition:  "Drosophila melanogaster muscle LIM protein at 84B (Mlp84B) gene, complete cds.",
	Accessions:  []string{"AF090832"},
	Version:     "AF090832.1",
	Keywords:    ".",
	Source:      "Drosophila melanogaster (fruit fly)",
	Organism:    "Drosophila melanogaster",
	OrganismTax: "Eukaryota; Metazoa; Ecdysozoa; Arthropoda; Hexapoda; Insecta; Pterygota; Neoptera; Endopterygota; Diptera; Brachycera; Muscomorpha; Ephydroidea; Drosophilidae; Drosophila; Sophophora.",
	References: []map[string]string{
		{
			"":        "1  (bases 1 to 5086)",
			"AUTHORS": "Stronach,B.E., Renfranz,P.J., Lilly,B. and Beckerle,M.C.",
			"TITLE":   "Muscle LIM proteins are associated with muscle sarcomeres and require dMEF2 for their expression during Drosophila myogenesis",
			"JOURNAL": "Mol. Biol. Cell 10 (7), 2329-2342 (1999)",
			"PUBMED":  "10397768",
		},
		{
			"":        "2  (bases 1 to 5086)",
			"AUTHORS": "Stronach,B.E., Renfranz,P.J., Lilly,B. and Beckerle,M.C.",
			"TITLE":   "Direct Submission",
			"JOURNAL": "Submitted (09-SEP-1998) Biology, University of Utah, 257 S. 1400 East, Salt Lake City, UT 84112-0840, USA",
		},
	},
	Features: []Feature{
		{
			Type: "source",
			Fields: map[string]string{
				"":           "1..5086",
				"organism":   "Drosophila melanogaster",
				"mol_type":   "genomic DNA",
				"db_xref":    "taxon:7227",
				"chromosome": "3",
				"map":        "3R; 84B-C",
			},
		},
		{
			Type: "gene",
			Fields: map[string]string{
				"":     "166..5010",
				"gene": "Mlp84B",
			},
		},
		{
			Type: "protein_bind",
			Fields: map[string]string{
				"":             "166..175",
				"gene":         "Mlp84B",
				"note":         "matches consensus at only 9 of 10 positions",
				"bound_moiety": "MEF2",
			},
		},
		{
			Type: "mRNA",
			Fields: map[string]string{
				"":     "join(1001..1055,2868..4655)",
				"gene": "Mlp84B",
			},
		},
		{
			Type: "5'UTR",
			Fields: map[string]string{
				"":     "join(1001..1055,2868..2946)",
				"gene": "Mlp84B",
			},
		},
		{
			Type: "protein_bind",
			Fields: map[string]string{
				"":             "1007..1016",
				"gene":         "Mlp84B",
				"note":         "site A",
				"bound_moiety": "MEF2",
			},
		},
		{
			Type: "protein_bind",
			Fields: map[string]string{
				"":             "1225..1234",
				"gene":         "Mlp84B",
				"note":         "site B",
				"bound_moiety": "MEF2",
			},
		},
		{
			Type: "protein_bind",
			Fields: map[string]string{
				"":             "1321..1330",
				"gene":         "Mlp84B",
				"note":         "site C",
				"bound_moiety": "MEF2",
			},
		},
		{
			Type: "protein_bind",
			Fields: map[string]string{
				"":             "1324..1333",
				"gene":         "Mlp84B",
				"experiment":   "experimental evidence, no additional details recorded",
				"note":         "site D; shown to bind Drosophila MEF2 in an in vitro gel shift assay",
				"bound_moiety": "MEF2",
			},
		},
		{
			Type: "protein_bind",
			Fields: map[string]string{
				"":             "1579..1588",
				"gene":         "Mlp84B",
				"note":         "site E",
				"bound_moiety": "MEF2",
			},
		},
		{
			Type: "protein_bind",
			Fields: map[string]string{
				"":             "2335..2344",
				"gene":         "Mlp84B",
				"note":         "matches consensus at only 9 of 10 positions",
				"bound_moiety": "MEF2",
			},
		},
		{
			Type: "CDS",
			Fields: map[string]string{
				"":            "2947..4434",
				"gene":        "Mlp84B",
				"note":        "MLP84B",
				"codon_start": "1",
				"product":     "muscle LIM protein at 84B",
				"protein_id":  "AAC61591.1",
				"translation": "MPSFQPIEAPKCPRCGKSVYAAEERLAGGYVFHKNCFKCGMCNK" +
					"SLDSTNCTEHERELYCKTCHGRKFGPKGYGFGTGAGTLSMDNGSQFLRENGDVPSVRN" +
					"GARLEPRAIARAPEGEGCPRCGGYVYAAEQMLARGRSWHKECFKCGTCKKGLDSILCC" +
					"EAPDKNIYCKGCYAKKFGPKGYGYGQGGGALQSDCYAHDDGAPQIRAAIDVDKIQARP" +
					"GEGCPRCGGVVYAAEQKLSKGREWHKKCFNCKDCHKTLDSINASDGPDRDVYCRTCYG" +
					"KKWGPHGYGFACGSGFLQTDGLTEDQISANRPFYNPDTTSIKARDGEGCPRCGGAVFA" +
					"AEQQLSKGKVWHKKCYNCADCHRPLDSVLACDGPDGDIHCRACYGKLFGPKGFGYGHA" +
					"PTLVSTSGESTIQFPDGRPLAGPKTSGGCPRCGFAVFAAEQMISKTRIWHKRCFYCSD" +
					"CRKSLDSTNLNDGPDGDIYCRACYGRNFGPKGVGYGLGAGALTTF",
			},
		},
		{
			Type: "3'UTR",
			Fields: map[string]string{
				"":     "4435..4655",
				"gene": "Mlp84B",
			},
		},
		{
			Type: "protein_bind",
			Fields: map[string]string{
				"":             "4629..4638",
				"gene":         "Mlp84B",
				"note":         "matches consensus at only 9 of 10 positions",
				"bound_moiety": "MEF2",
			},
		},
		{
			Type: "protein_bind",
			Fields: map[string]string{
				"":             "4955..4964",
				"gene":         "Mlp84B",
				"note":         "matches consensus at only 9 of 10 positions",
				"bound_moiety": "MEF2",
			},
		},
		{
			Type: "protein_bind",
			Fields: map[string]string{
				"":             "5001..5010",
				"gene":         "Mlp84B",
				"note":         "site F",
				"bound_moiety": "MEF2",
			},
		},
	},
	Origin: origin3,
}

// Taken from:
// https://www.ncbi.nlm.nih.gov/nucleotide/AF090832
const input3 = `LOCUS       AF090832                5086 bp    DNA     linear   INV 04-AUG-1999
DEFINITION  Drosophila melanogaster muscle LIM protein at 84B (Mlp84B) gene,
            complete cds.
ACCESSION   AF090832
VERSION     AF090832.1
KEYWORDS    .
SOURCE      Drosophila melanogaster (fruit fly)
  ORGANISM  Drosophila melanogaster
            Eukaryota; Metazoa; Ecdysozoa; Arthropoda; Hexapoda; Insecta;
            Pterygota; Neoptera; Endopterygota; Diptera; Brachycera;
            Muscomorpha; Ephydroidea; Drosophilidae; Drosophila; Sophophora.
REFERENCE   1  (bases 1 to 5086)
  AUTHORS   Stronach,B.E., Renfranz,P.J., Lilly,B. and Beckerle,M.C.
  TITLE     Muscle LIM proteins are associated with muscle sarcomeres and
            require dMEF2 for their expression during Drosophila myogenesis
  JOURNAL   Mol. Biol. Cell 10 (7), 2329-2342 (1999)
   PUBMED   10397768
REFERENCE   2  (bases 1 to 5086)
  AUTHORS   Stronach,B.E., Renfranz,P.J., Lilly,B. and Beckerle,M.C.
  TITLE     Direct Submission
  JOURNAL   Submitted (09-SEP-1998) Biology, University of Utah, 257 S. 1400
            East, Salt Lake City, UT 84112-0840, USA
FEATURES             Location/Qualifiers
     source          1..5086
                     /organism="Drosophila melanogaster"
                     /mol_type="genomic DNA"
                     /db_xref="taxon:7227"
                     /chromosome="3"
                     /map="3R; 84B-C"
     gene            166..5010
                     /gene="Mlp84B"
     protein_bind    166..175
                     /gene="Mlp84B"
                     /note="matches consensus at only 9 of 10 positions"
                     /bound_moiety="MEF2"
     mRNA            join(1001..1055,2868..4655)
                     /gene="Mlp84B"
     5'UTR           join(1001..1055,2868..2946)
                     /gene="Mlp84B"
     protein_bind    1007..1016
                     /gene="Mlp84B"
                     /note="site A"
                     /bound_moiety="MEF2"
     protein_bind    1225..1234
                     /gene="Mlp84B"
                     /note="site B"
                     /bound_moiety="MEF2"
     protein_bind    1321..1330
                     /gene="Mlp84B"
                     /note="site C"
                     /bound_moiety="MEF2"
     protein_bind    1324..1333
                     /gene="Mlp84B"
                     /experiment="experimental evidence, no additional details
                     recorded"
                     /note="site D; shown to bind Drosophila MEF2 in an in
                     vitro gel shift assay"
                     /bound_moiety="MEF2"
     protein_bind    1579..1588
                     /gene="Mlp84B"
                     /note="site E"
                     /bound_moiety="MEF2"
     protein_bind    2335..2344
                     /gene="Mlp84B"
                     /note="matches consensus at only 9 of 10 positions"
                     /bound_moiety="MEF2"
     CDS             2947..4434
                     /gene="Mlp84B"
                     /note="MLP84B"
                     /codon_start=1
                     /product="muscle LIM protein at 84B"
                     /protein_id="AAC61591.1"
                     /translation="MPSFQPIEAPKCPRCGKSVYAAEERLAGGYVFHKNCFKCGMCNK
                     SLDSTNCTEHERELYCKTCHGRKFGPKGYGFGTGAGTLSMDNGSQFLRENGDVPSVRN
                     GARLEPRAIARAPEGEGCPRCGGYVYAAEQMLARGRSWHKECFKCGTCKKGLDSILCC
                     EAPDKNIYCKGCYAKKFGPKGYGYGQGGGALQSDCYAHDDGAPQIRAAIDVDKIQARP
                     GEGCPRCGGVVYAAEQKLSKGREWHKKCFNCKDCHKTLDSINASDGPDRDVYCRTCYG
                     KKWGPHGYGFACGSGFLQTDGLTEDQISANRPFYNPDTTSIKARDGEGCPRCGGAVFA
                     AEQQLSKGKVWHKKCYNCADCHRPLDSVLACDGPDGDIHCRACYGKLFGPKGFGYGHA
                     PTLVSTSGESTIQFPDGRPLAGPKTSGGCPRCGFAVFAAEQMISKTRIWHKRCFYCSD
                     CRKSLDSTNLNDGPDGDIYCRACYGRNFGPKGVGYGLGAGALTTF"
     3'UTR           4435..4655
                     /gene="Mlp84B"
     protein_bind    4629..4638
                     /gene="Mlp84B"
                     /note="matches consensus at only 9 of 10 positions"
                     /bound_moiety="MEF2"
     protein_bind    4955..4964
                     /gene="Mlp84B"
                     /note="matches consensus at only 9 of 10 positions"
                     /bound_moiety="MEF2"
     protein_bind    5001..5010
                     /gene="Mlp84B"
                     /note="site F"
                     /bound_moiety="MEF2"
ORIGIN      
        1 tgatcaaacc taaagagtgg gacagagagt actactatat tcgtttcact cgccaaaagt
       61 tttgaacgca ggcgccttgc caatttttgt atcatactta taagtggaaa tgcaattgca
      121 aatgcatgtt aaaaatgttt tgagaaattc tgccgaaggt gcgcgttaat tatacaaaag
      181 gtgttgccca cccgttcttg tctttaagaa attccgtgtt tttcaatgct accccttgaa
      241 cttcggaatt ttcgatcgat ccatttttta aaatggtgtg tgtttcgatc gaaatactgt
      301 aataaatcaa atattaaatt agtttgttat aaatttgaga aacttactac tactactttg
      361 atttcaaaac taaaataatc ggtaaaatca taccattaga actaatatca aaatgaaatc
      421 attcttagta atgatattaa gaggtatttt ctggccgagt tggcccccgc catgtctggt
      481 tctattttgc cgatctgatt gatgctcaat gatgtacccc atgccccatt aatttccttt
      541 tgaaccaaaa atgcatggca aaatatcgaa ttttaattaa atgcaaatgc ggcatgttaa
      601 tttattgatg ttgaagtgcc aatgttccga aaggcaggca ggcagtggct ttgaacaaga
      661 cagttgaacc atgttgttgg tccaataaaa taccaacatc atctttttcg ccttgtcatg
      721 agcacgatca cctggaatat tctgctataa aatgaccgta tatacaattt atgaaagcac
      781 atttttcttg tatactagca tattgtatat gtgagancag atattttgtg atttgttggg
      841 attcttgttc ggatggctaa tgtatttatt gttggggcta tatagtcatc cctatcattt
      901 tctctggcga caaagttgcc ttcgtgccgt gtgcagacaa agactgaaag cgcttgacaa
      961 atttttcagc gcgatcagtc gaaaatcgtt gttaatcgaa gctaaactaa aattaaggat
     1021 atttctcttg cgaacctttt tattacgaac gaaaagtaaa tatcagtttt tgtcaatgga
     1081 aaacgaaaac caaaaccaaa accaatcacg cataacccag taaattccaa tattttgtgt
     1141 tgtgcttgtc taaaaaatat cggttcaagc tctgctacaa accaaacgcg aagcgatgaa
     1201 aatacagttt tgtcaagcga tcgtttaaat ttagaacttg tgcgtaattt ggaccataat
     1261 ttgacccaca aatgagagga ggaatgtatg gctccacaca ggaagcattt ttcccacaca
     1321 ctattattaa tagattccgt tttttcgagt gtgaaaaaac tttacggaaa ctgtctcact
     1381 caataaaatg agaactattt gcgaggagtt tttttttcaa tcgagtttta ttcaagtaca
     1441 aatcgatcat cttttcatca gatcaaatcg aatagtcact atccttgaat agagaaaaat
     1501 taaagataca gtccttactc gacgtaaaaa acacgtacat tttgatgagt tttataatct
     1561 tattgcatgc aaaaagtatt atatttagtc gggaatataa acaaactccc agatgcctga
     1621 ggtttccaat gatatgcaac ttttcgatca tctttgataa tctcccactg ggataaatgt
     1681 gcatataata tacacaaaag ttagcaaaca tccaagaatg agatacggtc aagaaaaaac
     1741 gagtgaagga tatggaataa aaaacttttt cgaagaatcg aaaaagttta aaaaccgcag
     1801 atgcatacca atatgtttcc atcccactct ctggcgttat ttgatgaaaa ctgaaataaa
     1861 acgaaggaaa actgctccgt gttgccgcat cttaagtttc tcaactcggt gtgctgagca
     1921 ggaaaactca gtggcagcgg cgaaaagtcg accgcgttgc cgtgccgtgt gcactcacaa
     1981 ttatatagac tagtttgctt tgagcaatct tccagttcag tactcggcgc ttccgtgtgt
     2041 ggttttagtg cgttgtaggt tccgttctat cttggctata atacactcac acccaacggt
     2101 tggatatcgc cagctcttaa ccgatttttt tgcacgcagc aaaagaccat tgcaaagcac
     2161 ctgactaatc catcgcaacc gctagataac cagattaccg gccagggacg atgctcatct
     2221 gatggtggtg ctttagctct gactcagttt tagaacaccc agatatgtac atatatatat
     2281 ctgccacagt tattttttgt gtgtgctttt cccaaatcta tctacctgtc tgcgttattt
     2341 ttatacaaca ttttctgtat agatagatac acatttattg ctcatgtttt tctctgtgtt
     2401 ctgtgatttg aaaaattcct cctttatttt ttgttgtcat gagaaagatg ttattttctg
     2461 atcattaaca taggaatctt caatatgaat tacgtataca tacgtccggt tgggtaattt
     2521 tctggcacgg gcgatgacgt ctcgactgac agataagaat cttgatagat gatctacata
     2581 atgcaaattg aggagaaatt cgcgactgtc aagttgaatt tgtaatagtt cattgtaaac
     2641 ttaagataaa acaagcgtct taatcgtcca aataggttaa catgtagatt agcttgaatt
     2701 aactagttgg gcttagaata atttggaaac catttaacat acaaaaatga cagatgctta
     2761 cctcatcgct ttcatcgcca ccaaagcttc cactctgctg catttcccag caggacttac
     2821 gccacgatcc catcgactct atgtgttctt ctattttgct atttcagata ataaacttaa
     2881 taagtaaagt aattgtactc gcgtagatca ttttgatagc gttaaaaaaa ccgaaacaca
     2941 gacagaatgc cttccttcca accgattgag gcccccaagt gtccgcgctg cggcaagagt
     3001 gtctacgctg ccgaagagcg tctggctggc ggctatgtat tccacaagaa ctgcttcaag
     3061 tgcggaatgt gcaacaaatc cctggactcc accaactgca cagagcacga gcgcgagctc
     3121 tactgcaaga cgtgccacgg tcgcaagttc gggccgaaag gttacggctt cggcactgga
     3181 gctggcaccc tctccatgga caacgggtca cagttcctgc gcgagaacgg cgatgtgccg
     3241 tccgtaagga atggagcccg cctggaaccc agggctattg ctcgtgcccc cgaaggtgag
     3301 ggctgtcctc gttgcggtgg ctatgtgtac gccgccgaac agatgttagc ccgcggtcgc
     3361 agctggcaca aggagtgctt caagtgcggt acctgcaaga agggtctgga ctcgatcctg
     3421 tgctgcgagg ctccggacaa gaacatctac tgcaagggct gctatgccaa gaagtttgga
     3481 cccaagggct atggttatgg ccagggcggt ggtgctctcc agtccgactg ctatgctcac
     3541 gacgacggag caccgcaaat ccgtgccgcc attgatgtgg acaagatcca ggcccgtccg
     3601 ggtgagggtt gcccacgttg cggcggtgtg gtctacgcag cggagcagaa gctttccaag
     3661 ggtcgggagt ggcacaagaa gtgcttcaac tgcaaggatt gccacaagac tctggactcg
     3721 atcaatgcca gygatggtcc cgatcgtgat gtgtactgcc gcacctgcta cggcaagaag
     3781 tggggaccac atggctatgg attcgcatgc ggctctggtt tcctgcagac cgatggcttg
     3841 accgaggatc agatcagcgc caacaggccc ttctataacc cggacaccac gtcaattaag
     3901 gcccgtgacg gcgaaggctg tccccggtgc ggaggagccg tattcgccgc cgagcaacag
     3961 ctgtccaagg gcaaggtgtg gcacaagaag tgctacaact gcgccgactg ccaccggcca
     4021 ttggactcgg tcctggcctg cgatggaccc gatggcgaca tccactgccg cgcctgctac
     4081 ggcaagctct tcggccccaa gggctttggt tacggccacg cccccactct agtgtcaacc
     4141 agtggcgaga gcaccatcca gttcccagat ggccgtcctc tggccggacc caagacttcg
     4201 ggcggctgcc cgcgttgcgg tttcgccgtg ttcgccgccg agcagatgat cagcaagacc
     4261 aggatctggc acaagaggtg cttctactgc tcggattgcc gcaaatcgct ggactcgacc
     4321 aacctgaacg atggacccga cggcgatatc tactgccgag cctgctacgg ccgcaatttt
     4381 ggacccaagg gagtgggcta cggtctgggc gcgggcgctt tgacaacgtt ctaaacgagc
     4441 ctattttatg tatatccaat cgaaacccac atccatgtct gtgaaccgtt ccgatgcatc
     4501 tgtctaatga ccaatctgtt ttaactaact atccctgaaa ttagctaact tagtcttcgt
     4561 tttgcttcca acgatttttg atcaacgatg cactgaatgc aattcaattt tctttaggtt
     4621 atgtgagatt aagtttaata tgaaaacaag aaaaaaagaa taaatgccaa aaacataaaa
     4681 cttacgactc aaagcgaaac taaccatcga acacagacta acgaactaaa cgacattgac
     4741 cacatcgaga aaagcaaaac tttctgcagg tgaaactgct gcgagccaaa gtcgttatta
     4801 cttttttgta tcttaaacta tttattttcg aagaccagcc aagtatgtct tatgctccaa
     4861 atcgaattac ccaaaagcat actttgtgca aatggaaaaa ttttatacac ataaatatac
     4921 tcttacgacg aacaatgatg aaagatatga tgttgtatat ttagctgatt tagttcttac
     4981 gatatactta tgaaatctga ttatttataa tgatgtttga atggatttct cacccatcta
     5041 atctcttgca ttctgttttg gggatgccag ccatgggcag atatta
//`

const origin3 = "tgatcaaacctaaagagtgggacagagagtactactatattcgtttcactcgccaaaagt" +
	"tttgaacgcaggcgccttgccaatttttgtatcatacttataagtggaaatgcaattgca" +
	"aatgcatgttaaaaatgttttgagaaattctgccgaaggtgcgcgttaattatacaaaag" +
	"gtgttgcccacccgttcttgtctttaagaaattccgtgtttttcaatgctaccccttgaa" +
	"cttcggaattttcgatcgatccattttttaaaatggtgtgtgtttcgatcgaaatactgt" +
	"aataaatcaaatattaaattagtttgttataaatttgagaaacttactactactactttg" +
	"atttcaaaactaaaataatcggtaaaatcataccattagaactaatatcaaaatgaaatc" +
	"attcttagtaatgatattaagaggtattttctggccgagttggcccccgccatgtctggt" +
	"tctattttgccgatctgattgatgctcaatgatgtaccccatgccccattaatttccttt" +
	"tgaaccaaaaatgcatggcaaaatatcgaattttaattaaatgcaaatgcggcatgttaa" +
	"tttattgatgttgaagtgccaatgttccgaaaggcaggcaggcagtggctttgaacaaga" +
	"cagttgaaccatgttgttggtccaataaaataccaacatcatctttttcgccttgtcatg" +
	"agcacgatcacctggaatattctgctataaaatgaccgtatatacaatttatgaaagcac" +
	"atttttcttgtatactagcatattgtatatgtgagancagatattttgtgatttgttggg" +
	"attcttgttcggatggctaatgtatttattgttggggctatatagtcatccctatcattt" +
	"tctctggcgacaaagttgccttcgtgccgtgtgcagacaaagactgaaagcgcttgacaa" +
	"atttttcagcgcgatcagtcgaaaatcgttgttaatcgaagctaaactaaaattaaggat" +
	"atttctcttgcgaacctttttattacgaacgaaaagtaaatatcagtttttgtcaatgga" +
	"aaacgaaaaccaaaaccaaaaccaatcacgcataacccagtaaattccaatattttgtgt" +
	"tgtgcttgtctaaaaaatatcggttcaagctctgctacaaaccaaacgcgaagcgatgaa" +
	"aatacagttttgtcaagcgatcgtttaaatttagaacttgtgcgtaatttggaccataat" +
	"ttgacccacaaatgagaggaggaatgtatggctccacacaggaagcatttttcccacaca" +
	"ctattattaatagattccgttttttcgagtgtgaaaaaactttacggaaactgtctcact" +
	"caataaaatgagaactatttgcgaggagtttttttttcaatcgagttttattcaagtaca" +
	"aatcgatcatcttttcatcagatcaaatcgaatagtcactatccttgaatagagaaaaat" +
	"taaagatacagtccttactcgacgtaaaaaacacgtacattttgatgagttttataatct" +
	"tattgcatgcaaaaagtattatatttagtcgggaatataaacaaactcccagatgcctga" +
	"ggtttccaatgatatgcaacttttcgatcatctttgataatctcccactgggataaatgt" +
	"gcatataatatacacaaaagttagcaaacatccaagaatgagatacggtcaagaaaaaac" +
	"gagtgaaggatatggaataaaaaactttttcgaagaatcgaaaaagtttaaaaaccgcag" +
	"atgcataccaatatgtttccatcccactctctggcgttatttgatgaaaactgaaataaa" +
	"acgaaggaaaactgctccgtgttgccgcatcttaagtttctcaactcggtgtgctgagca" +
	"ggaaaactcagtggcagcggcgaaaagtcgaccgcgttgccgtgccgtgtgcactcacaa" +
	"ttatatagactagtttgctttgagcaatcttccagttcagtactcggcgcttccgtgtgt" +
	"ggttttagtgcgttgtaggttccgttctatcttggctataatacactcacacccaacggt" +
	"tggatatcgccagctcttaaccgatttttttgcacgcagcaaaagaccattgcaaagcac" +
	"ctgactaatccatcgcaaccgctagataaccagattaccggccagggacgatgctcatct" +
	"gatggtggtgctttagctctgactcagttttagaacacccagatatgtacatatatatat" +
	"ctgccacagttattttttgtgtgtgcttttcccaaatctatctacctgtctgcgttattt" +
	"ttatacaacattttctgtatagatagatacacatttattgctcatgtttttctctgtgtt" +
	"ctgtgatttgaaaaattcctcctttattttttgttgtcatgagaaagatgttattttctg" +
	"atcattaacataggaatcttcaatatgaattacgtatacatacgtccggttgggtaattt" +
	"tctggcacgggcgatgacgtctcgactgacagataagaatcttgatagatgatctacata" +
	"atgcaaattgaggagaaattcgcgactgtcaagttgaatttgtaatagttcattgtaaac" +
	"ttaagataaaacaagcgtcttaatcgtccaaataggttaacatgtagattagcttgaatt" +
	"aactagttgggcttagaataatttggaaaccatttaacatacaaaaatgacagatgctta" +
	"cctcatcgctttcatcgccaccaaagcttccactctgctgcatttcccagcaggacttac" +
	"gccacgatcccatcgactctatgtgttcttctattttgctatttcagataataaacttaa" +
	"taagtaaagtaattgtactcgcgtagatcattttgatagcgttaaaaaaaccgaaacaca" +
	"gacagaatgccttccttccaaccgattgaggcccccaagtgtccgcgctgcggcaagagt" +
	"gtctacgctgccgaagagcgtctggctggcggctatgtattccacaagaactgcttcaag" +
	"tgcggaatgtgcaacaaatccctggactccaccaactgcacagagcacgagcgcgagctc" +
	"tactgcaagacgtgccacggtcgcaagttcgggccgaaaggttacggcttcggcactgga" +
	"gctggcaccctctccatggacaacgggtcacagttcctgcgcgagaacggcgatgtgccg" +
	"tccgtaaggaatggagcccgcctggaacccagggctattgctcgtgcccccgaaggtgag" +
	"ggctgtcctcgttgcggtggctatgtgtacgccgccgaacagatgttagcccgcggtcgc" +
	"agctggcacaaggagtgcttcaagtgcggtacctgcaagaagggtctggactcgatcctg" +
	"tgctgcgaggctccggacaagaacatctactgcaagggctgctatgccaagaagtttgga" +
	"cccaagggctatggttatggccagggcggtggtgctctccagtccgactgctatgctcac" +
	"gacgacggagcaccgcaaatccgtgccgccattgatgtggacaagatccaggcccgtccg" +
	"ggtgagggttgcccacgttgcggcggtgtggtctacgcagcggagcagaagctttccaag" +
	"ggtcgggagtggcacaagaagtgcttcaactgcaaggattgccacaagactctggactcg" +
	"atcaatgccagygatggtcccgatcgtgatgtgtactgccgcacctgctacggcaagaag" +
	"tggggaccacatggctatggattcgcatgcggctctggtttcctgcagaccgatggcttg" +
	"accgaggatcagatcagcgccaacaggcccttctataacccggacaccacgtcaattaag" +
	"gcccgtgacggcgaaggctgtccccggtgcggaggagccgtattcgccgccgagcaacag" +
	"ctgtccaagggcaaggtgtggcacaagaagtgctacaactgcgccgactgccaccggcca" +
	"ttggactcggtcctggcctgcgatggacccgatggcgacatccactgccgcgcctgctac" +
	"ggcaagctcttcggccccaagggctttggttacggccacgcccccactctagtgtcaacc" +
	"agtggcgagagcaccatccagttcccagatggccgtcctctggccggacccaagacttcg" +
	"ggcggctgcccgcgttgcggtttcgccgtgttcgccgccgagcagatgatcagcaagacc" +
	"aggatctggcacaagaggtgcttctactgctcggattgccgcaaatcgctggactcgacc" +
	"aacctgaacgatggacccgacggcgatatctactgccgagcctgctacggccgcaatttt" +
	"ggacccaagggagtgggctacggtctgggcgcgggcgctttgacaacgttctaaacgagc" +
	"ctattttatgtatatccaatcgaaacccacatccatgtctgtgaaccgttccgatgcatc" +
	"tgtctaatgaccaatctgttttaactaactatccctgaaattagctaacttagtcttcgt" +
	"tttgcttccaacgatttttgatcaacgatgcactgaatgcaattcaattttctttaggtt" +
	"atgtgagattaagtttaatatgaaaacaagaaaaaaagaataaatgccaaaaacataaaa" +
	"cttacgactcaaagcgaaactaaccatcgaacacagactaacgaactaaacgacattgac" +
	"cacatcgagaaaagcaaaactttctgcaggtgaaactgctgcgagccaaagtcgttatta" +
	"cttttttgtatcttaaactatttattttcgaagaccagccaagtatgtcttatgctccaa" +
	"atcgaattacccaaaagcatactttgtgcaaatggaaaaattttatacacataaatatac" +
	"tcttacgacgaacaatgatgaaagatatgatgttgtatatttagctgatttagttcttac" +
	"gatatacttatgaaatctgattatttataatgatgtttgaatggatttctcacccatcta" +
	"atctcttgcattctgttttggggatgccagccatgggcagatatta"
