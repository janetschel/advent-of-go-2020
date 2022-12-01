day = $(shell date +'%-d')
year = $(shell date +'%-Y')

date = $(shell printf "%02d" $(day))

new:
	@echo "Creating new file structure for day" $(date) $(year)"..."
	
	mkdir -p calendar/$(year); \
	mkdir -p calendar/$(year)/day-$(date); \
	cp template calendar/$(year)/day-$(date)/day$(date).go; \
	ex +%s/inputDay/$(day)/g -scwq calendar/$(year)/day-$(date)/day$(date).go; \
	ex +%s/inputYear/$(year)/g -scwq calendar/$(year)/day-$(date)/day$(date).go; \
	echo "Files successfully created.. happy hacking :)"
	@echo "INFO: puzzle input still needs to be fetched"
	@git add calendar/

all:
	number=1 ; while [ "$$number" -le 25 ] ; do \
			$(MAKE) new day=$$number year=$$year ; \
			number=$$(( $$number + 1)); \
		done
