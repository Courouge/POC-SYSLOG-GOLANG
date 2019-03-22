L'application est codé en scala (version 2.11).

Les middlewares utilisés sont: Kafka -> Spark Streaming -> Vertica

La gestion des offsets est géré par le consumer group dans Kafka.

Les messages qui comportent des erreurs sont stockés dans un topic error

Les messages qui ont une structure json avec des champs supplementaires sont injectés dans la table vertica, les champs supplementaires ne sont pas sauvegardés.

Problemes rencontrés:
- L'injection de données dans Kafka contenant des carateres speciaux(',").
- Le driver Vertica ne gere pas les injections de donnée null pour les dates.
- La base de donnée Vertica ne gere pas l'injection de date bidon comme 0000-00-00.
