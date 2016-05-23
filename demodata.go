//
// Copyright (c) 2016 Dean Jackson <deanishe@deanishe.net>
//
// MIT Licence. See http://opensource.org/licenses/MIT
//
// Created on 2016-05-23
//

package main

// Useful for screenshots
var testHostnames = []string{
	"wiek.segarra-pavon.com",
	"bustamante.romano-leone.com",
	"kozey.brakus.info",
	"ruggiero-ferrara.piras.info",
	"greco-riva.figuerola.com",
	"manzano-tomas.dangelo.com",
	"ledesma.arana-pinol.com",
	"valle.moen.com",
	"costantini.giuliani.org",
	"schuppe.morcillo.com",
	"sanmiguel-ferreras.powlowski-stracke.net",
	"watsica-lang.ferraro.com",
	"barrows.feeney.com",
	"kostolzin.rosenbaum-kulas.biz",
	"carnero.henk.com",
	"schmidt.cremin-moore.com",
	"romano.salamanca-cervera.com",
	"carbonell.oliver-pozo.com",
	"keudel.bianchi.com",
	"russo-rossetti.moya-leon.com",
	"grasso.huebel.de",
	"serra.wesack.com",
	"rowe.giordano-fabbri.info",
	"weitzel.eichmann.com",
	"reising.giordano-montanari.com",
	"piras-lombardi.martini.com",
	"vandervort-powlowski.jodar.com",
	"ohara.kilback-wisozk.org",
	"kuphal.barone.com",
	"kallert.gonzalez.com",
	"bien.sobrino.com",
	"mies.stracke-bogan.biz",
	"caputo.johnston.com",
	"weiss.gulgowski.com",
	"sanford-keeling.martino.com",
	"colombo-marino.klocko-bernier.info",
	"austermuehle.kautzer.org",
	"saeuberlich.sartori.org",
	"fischer.cerda-alvarez.com",
	"marquez.knappe.com",
	"bruder.benedetti-dangelo.org",
	"caputo.marini.com",
	"conti.valentini.com",
	"salvador.bruen-kerluke.net",
	"rath.roman-barral.com",
	"narvaez.santoro.biz",
	"gertz.wisoky-harvey.com",
	"ferretti.suarez.com",
	"klapp.trantow.info",
	"trubin.cremin.org",
	"junitz.bartell-schaden.com",
	"battaglia.villena.com",
	"monti-orlando.lehner-orn.com",
	"valentini.pellegrini.net",
	"almansa.ferrara.com",
	"arteaga-molina.jopich.de",
	"cattaneo.giuliani.com",
	"piras.guerra.biz",
	"cartwright.piras.info",
	"ocana.huebel.org",
	"scheuermann.saenz.info",
	"buckridge-maggio.marquez.com",
	"abascal-alba.russo-messina.com",
	"heuser.buchholz.net",
	"de.testa.com",
	"costa-parisi.sucker.de",
	"radisch.ferraro.com",
	"valenzuela-reyes.paffrath.com",
	"beer.neureuther.org",
	"romano.geisel.com",
	"herrmann.haering.net",
	"blanda.heintze.com",
	"heinrich.zboncak.biz",
	"neri-sartori.mende.org",
	"sala-de.roma.com",
	"capdevila.monti-conti.biz",
	"rossi.damico.com",
	"geisler.grimes.com",
	"uriarte.ferretti.com",
	"king.benedetti-bernardi.biz",
	"hahn.schueler.com",
	"vidal.rogahn-skiles.com",
	"de.saeuberlich.com",
	"rinaldi.moretti.com",
	"runte-oconnell.conte.org",
	"grady.santoro.com",
	"heidenreich.bellini.info",
	"bruno.wiza.net",
	"salazar-cordoba.franco.com",
	"martino-mazza.montanari-ferraro.org",
	"denesik.kling.com",
	"abshire-hilpert.hethur.com",
	"sorgatz.borrego.com",
	"barrera.mayert.com",
	"legros.toy.biz",
	"geissler.ferraro.com",
	"hahn.colombo.org",
	"lowe.mraz.info",
	"kuhl.leannon-jewess.net",
	"noack.montoya.info",
	"reichmann.almeida-cortes.com",
	"klocko-leuschke.milani-mazza.com",
	"falco-tamarit.speer.org",
	"vitali.zulauf-dare.net",
	"borras-valentin.feest-kilback.com",
	"reichmann.dippel.com",
	"hoffmann.conroy.org",
	"blanca.hermighausen.org",
	"marino-rossetti.valencia-salamanca.net",
	"heintze.angel-agudo.com",
	"palmieri.swaniawski-konopelski.com",
	"raedel.tintzmann.de",
	"ruggiero-farina.huguet.com",
	"padberg.losa-fuster.com",
	"blasco.gotthard.com",
	"speer.davids.com",
	"caputo-moretti.jacobi-rau.com",
	"freudenberger.granados.com",
	"bosco.armstrong.com",
	"sorrentino.haase.net",
	"parisian-hauck.connelly.com",
	"mosemann.koch-green.com",
	"benavides.palumbo-costa.com",
	"roldan-conesa.larkin-white.biz",
	"deckow.pergande.com",
	"revilla-fajardo.martini.com",
	"carreno-bartolome.villalonga.com",
	"watsica.grady.info",
	"mora-conde.berge.net",
	"sartori.heras.com",
	"mielcarek.mayo-galvan.com",
	"henschel.muehle.de",
	"mude.kranz.com",
	"rovira-ramirez.sastre.com",
	"ruecker.bashirian-koelpin.org",
	"gutmann.hettner.com",
	"effertz-emmerich.trapp.com",
	"herman.barba-vara.info",
	"boyer.saez.com",
	"blazquez.peinado.com",
	"durgan.weitzel.net",
	"figuerola-duran.martini.net",
	"crona.bahringer-cremin.net",
	"martinez.white.com",
	"bruno.damico.org",
	"arce-adan.scheibe.de",
	"parisi-palmieri.bustos-corbacho.biz",
	"bien.pellegrino.net",
	"okuneva-moen.barth.com",
	"santamaria-arias.schmidt.de",
	"bonbach.ullmann.com",
	"ratke.lorenzo-torrents.com",
	"bachmann.schumm.info",
	"trujillo-molina.mariani-russo.com",
	"mariani-carbone.caruso.biz",
	"santamaria-medina.cabrero.biz",
	"cuellar.reig-figuerola.info",
	"dietrich.trupp.net",
	"hettner.ritchie-rowe.net",
	"farina.reichert.com",
	"ferraro.arjona.net",
	"martinelli-de.gude.com",
	"rippin-mitchell.mans.org",
	"donati.rossi.info",
	"malo.robles.biz",
	"schottin.lobato.com",
	"zaenker.yanez.org",
	"barrena.cozar.com",
	"nienow.sala.com",
	"casper.bellini-rinaldi.com",
	"pfeffer.robledo.com",
	"hartung.ramirez-barragan.net",
	"pieper.beier.org",
	"spiess.ruggiero.com",
	"monti.rizzo.com",
	"infante-baron.rueda.com",
	"segura-yanez.villar.com",
	"stamm.robel.org",
	"abad-velez.quitzon.com",
	"gallo.canizares.com",
	"satterfield.heathcote.com",
	"ernst.miro.com",
	"nader.hernando.com",
	"matas.riquelme-vigil.com",
	"gatti-amato.dussen.com",
	"haenel.martini-barone.com",
	"hendriks.vigil.com",
	"finke.pinedo-egea.org",
	"catalan.ferretti.com",
	"saenz.montero.com",
	"seifert.shanahan.org",
	"glover.butte.de",
	"ruppersberger.benedetti-rossetti.com",
	"weller.vitale.com",
	"penas.sevillano.net",
	"bruno-villa.thompson-dare.com",
	"pla.trubin.de",
	"castilla-bonilla.mcglynn.com",
	"ledner.baerer.org",
	"bruno.barone.com",
}