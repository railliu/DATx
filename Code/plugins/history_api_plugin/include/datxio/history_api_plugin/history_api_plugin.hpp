/**
 *  @file
 *  @copyright defined in datx/LICENSE.txt
 */

#pragma once
#include <datxio/history_plugin/history_plugin.hpp>
#include <datxio/core_plugin/core_plugin.hpp>
#include <datxio/http_plugin/http_plugin.hpp>

#include <appbase/application.hpp>

namespace datxio {

   using namespace appbase;

   class history_api_plugin : public plugin<history_api_plugin> {
      public:
        APPBASE_PLUGIN_REQUIRES((history_plugin)(core_plugin)(http_plugin))

        history_api_plugin();
        virtual ~history_api_plugin();

        virtual void set_program_options(options_description&, options_description&) override;

        void plugin_initialize(const variables_map&);
        void plugin_startup();
        void plugin_shutdown();

      private:
   };

}
