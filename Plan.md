We'll build this out in a tick/tock fashion
Ticks will be infrastructure improvements, Tocks will be feature driven 

#### v0.0 - SET THINGS UP RIGHT
Requirements:
    - Service Discovery
    - Health
    - Logging/Monitoring
    - Debugging Production
    - CI (Stage/Prod)

#### v0.1 - IT HAS TO ACTUALLY BE A THING
Features:
    - Users can create bots
    - Users can delete bots
    - Users can configure bot personality
    - Bots can post status updates
        - Constraint: They can only spam randomly generated messages
    - Users can view a bot's profile and see all of the status updates posted
    - Bots can follow other bots
        - Constraint: They are given a random bot every day If the bot is configured to follow bots, they will automatically follow them 
    - Bots can comment on a status update
        - Constraint: They can only spam randomly generated message
        
Requirements:
    - Tests
    - Feature toggles
    
#### v0.2 - MAKE IT PERFORMANT
Requirements:
    - Kafka
    - Samza
    - Seamless transition Users should experience 0 downtime
    
#### v0.x - TBD

#### v1.0 - Expose it to the world!
Requirements:
    - Load testing
    - Easily scalable
    
#### v2.0 - ????
Requirements:
    - VC funding
    - Mission/Vision Statement
    - Swank office space

#### v3.0 - PROFIT
    - IPO