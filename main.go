package main

import "log"

const (
	v       int    = 23
	g       int    = 92
	keyLen         = 10 + (v+g)%7
	message string = `At an abandoned hotel, a police squad corners Trinity, who overpowers them with superhuman abilities. She flees, pursued by the police and a group of suited Agents capable of similar superhuman feats. She answers a ringing public telephone and vanishes.

	Computer programmer Thomas Anderson, known by his hacking alias "Neo", is puzzled by repeated online encounters with the phrase "the Matrix". Trinity contacts him and tells him a man named Morpheus has the answers Neo seeks. A team of Agents and police, led by Agent Smith, arrives at Neo's workplace in search of him. Though Morpheus attempts to guide Neo to safety, he surrenders rather than risk a dangerous escape via a scaffold. The Agents attempt to bribe Neo into helping them locate Morpheus, who they claim is a terrorist (and the most dangerous man alive), by offering to erase his criminal record. When he refuses to cooperate, they fuse his mouth shut, pin him down, and implant a robotic "bug" in his stomach. Neo wakes up from what he believes to be a nightmare. Soon after, Neo is taken by Trinity to meet Morpheus, and she removes the bug from him, indicating that the "nightmare" he experienced was apparently real.
	
	Morpheus offers Neo a choice between two pills: red to reveal the truth about the Matrix, and blue to forget everything and return to his former life. As Neo takes the red pill, his reality begins to distort, and he soon awakens in a liquid-filled pod among countless other pods, containing other humans. He is then brought aboard Morpheus's flying ship, the Nebuchadnezzar. As Neo recuperates from a lifetime of physical inactivity in the pod in the aftermath of being awakened, Morpheus explains the situation: In the early 21st century, a war broke out between humanity and intelligent machines. After humans blocked the machines' access to solar energy by using nuclear weapons, the machines responded by capturing humans and harvesting their bioelectric power, while keeping their minds pacified in the Matrix, a shared simulated reality modeled on the world as it was in 1999. The machines won the war and the remaining free humans took refuge in the underground city of Zion.
	
	Morpheus and his crew are a group of rebels who hack into the Matrix to "unplug" enslaved humans and recruit them; their understanding of the Matrix's simulated nature allows them to bend its physical laws. Morpheus warns Neo that death within the Matrix kills the physical body too and explains that the Agents are sentient programs that eliminate threats to the system, while machines called Sentinels eliminate rebels in the real world. Neo's prowess during virtual training cements Morpheus's belief that Neo is "the One", a human prophesied to free humankind. The group enters the Matrix to visit the Oracle, the prophet who predicted that the One would emerge. She implies to Neo that he is not the One and warns that he will have to choose between Morpheus's life and his own. Before they can leave the Matrix, Agents and police ambush the group, tipped off by Cypher, a disgruntled crew member who has betrayed Morpheus in exchange for a comfortable life in the Matrix.
	
	To buy time for the others, Morpheus fights Smith and is captured. Cypher exits the Matrix and murders the other crew members as they lie unconscious. Before Cypher can kill Neo and Trinity, crew member Tank regains consciousness and kills him before pulling Neo and Trinity from the Matrix. The Agents interrogate Morpheus to learn his access codes to the mainframe computer in Zion, which would allow them to destroy it. Neo resolves to return to the Matrix to rescue Morpheus, as the Oracle prophesied; Trinity insists she accompany him. While rescuing Morpheus, Neo gains confidence in his abilities, performing feats comparable to those of the Agents.
	
	After Morpheus and Trinity safely exit the Matrix, Smith ambushes and kills Neo. While a group of Sentinels attack the Nebuchadnezzar, Trinity confesses her love for Neo and says the Oracle told her she would fall in love with the One. Neo is revived, with newfound abilities to perceive and control the Matrix; he easily defeats Smith, prompting the other Agents to flee and he leaves the Matrix just as the ship's electromagnetic pulse disables the Sentinels. Back in the Matrix, Neo makes a telephone call, promising the machines that he will show their prisoners "a world where anything is possible". He hangs up and flies away.`
)

func main() {
	server, err := NewServer()
	if err != nil {
		log.Fatalln("error creating server", err)
	}
	client := NewClient()
	client.ConnectTo(server)
	go func() {
		err := server.Boot()
		if err != nil {
			log.Fatalln(err)
		}
	}()
	err = client.Work()
	if err != nil {
		log.Fatalln(err)
	}
}
